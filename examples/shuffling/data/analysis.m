berger =[4 4 3 2 216 198 241;
    5 3 3 2 155 160 172;
    5 4 3 1 129 116 124;
    5 4 2 2 106 92 105;
    4 3 3 3 105 103 129;
    6 3 2 2 56 64 46;
    6 4 2 1 47 53 36;
    6 3 3 1 34 40 41;
    5 5 2 1 32 40 19;
    4 4 4 1 30 35 25;
    7 3 2 1 90 99 62]

shuffles = [];
a=importdata('jessie.shuffles');
for i = 1:size(a,1)
    for j=1:size(a,2)
        if isnan(a(i,j))==0
            shuffles = [shuffles; a(i,j)];
        end
    end
end
a=importdata('zack.shuffles');
for i = 1:size(a,1)
    for j=1:size(a,2)
        if isnan(a(i,j))==0
            shuffles = [shuffles; a(i,j)];
        end
    end
end

[n] = hist(shuffles,1:5)
cdfShuffles = cumsum(n)/sum(n);
plot(cdfShuffles)

test = [];
for i=1:1000
    r = rand;
    for j=1:length(cdfShuffles)
        if r < cdfShuffles(j)
            break
        end
    end
    test = [test; j];
end
c{1}=shuffles;
c{2}=test;
nhist(c,'samebins',1)

deck = [1:13 101:113 1001:1013 10001:10013];
bergerTest = zeros(size(berger,1),1);

for iterations=1:1000
    deck = deck(randperm(length(deck)));
    suits = [0 0 0 0];
    for i=1:13
        if deck(i) > 0 && deck(i) < 100
            suits(1) = suits(1)+1;
        elseif deck(i) > 100 && deck(i) < 1000
            suits(2) = suits(2)+1;
        elseif deck(i) > 1000 && deck(i) < 10000
            suits(3) = suits(3)+1;
        elseif deck(i) > 10000 && deck(i) < 100000
            suits(4) = suits(4)+1;
        end
    end
    suits = fliplr(sort(suits));

    for i=1:size(berger,1)
        for j=1:4
            if suits(j)~=berger(i,j)
                break
            end
        end
        if j==4
            break
        end
    end
    bergerTest(i)= bergerTest(i)+1;
end

hold off;
for numShuffles=1:100
    myTest = zeros(size(berger,1),1);
    for iterations=1:1000
        setdeck = [1:13 101:113 1001:1013 10001:10013];
        for shuffles=1:numShuffles
            % Generate the packets
            packets = [];
            topDeck = 0;
            for i=1:52
                r = rand;
                for j=1:length(cdfShuffles)
                    if r < cdfShuffles(j)
                        break
                    end
                end
                if sum([packets; j]) > 52
                    packets = [packets; 52 - sum(packets)];
                else
                    packets = [packets; j];
                end
                if mod(i,2)==1
                    topDeck = topDeck + packets(i);
                end
                if sum(packets)==52
                    break
                end
            end

            % Organize the packets according to the shuffle
            packet1=1:topDeck;
            packet2=topDeck+1:52;
            if rand < 0.5
                packet1=52-topDeck+1:52;
                packet2=1:52-topDeck;
            end
            packet1i = 0;
            packet2i = 0;
            cardNums = zeros(52,1);
            cardi = 0;
            for i=1:length(packets)
                if mod(i,2)==1
                    for j=1:packets(i)
                        cardi = cardi + 1;
                        packet1i = packet1i + 1;
                        cardNums(cardi) = packet1(packet1i);
                    end
                else
                    for j=1:packets(i)
                        cardi = cardi + 1;
                        packet2i = packet2i + 1;
                        cardNums(cardi) = packet2(packet2i);
                    end            
                end
            end

            % Redraw the deck
            deck = zeros(52,1);
            for i=1:52
                deck(i)=setdeck(cardNums(i));
            end
            setdeck = deck;
        end

        suits = [0 0 0 0];
        for i=1:13
            if deck(i) > 0 && deck(i) < 100
                suits(1) = suits(1)+1;
            elseif deck(i) > 100 && deck(i) < 1000
                suits(2) = suits(2)+1;
            elseif deck(i) > 1000 && deck(i) < 10000
                suits(3) = suits(3)+1;
            elseif deck(i) > 10000 && deck(i) < 100000
                suits(4) = suits(4)+1;
            end
        end
        suits = fliplr(sort(suits));

        for i=1:size(berger,1)
            for j=1:4
                if suits(j)~=berger(i,j)
                    break
                end
            end
            if j==4
                break
            end
        end
        myTest(i)= myTest(i)+1;
    %     if i==11
    %         disp(suits)
    %         break
    %     end
    end
plot(myTest/sum(myTest))
hold on;
end
plot(bergerTest/sum(bergerTest))
plot(berger(:,5)/sum(berger(:,5)))
plot(berger(:,6)/sum(berger(:,6)))
plot(berger(:,7)/sum(berger(:,7)))
legend('1 shuffle','2 shuffles','3 shuffles','random simulated','berger - expected','berger - actual comp','berger - actual man')
xlabel('Kind of hand (Berger et al.)')
ylabel('Probability')
title('Distribution of suits in 13-card hand')
