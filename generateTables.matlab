dfs = 1:1000;
Pvalues = [0.9999 0.999 0.995	0.975	0.20	0.10	0.05	0.025	0.02	0.01	0.005	0.002	0.001 0.0001];
Pvalues = [Pvalues 0.01:0.01:0.99];
Pvalues = sort(Pvalues);
Pvalues = unique(Pvalues);

dfs = 1:2;
Pvalues = [0.9999 0.999 0.995	0.975	0.20	0.10	0.05	0.025	0.02	0.01	0.005	0.002	0.001 0.0001];
Pvalues = sort(Pvalues);
Pvalues = unique(Pvalues);
chiSquareMatrix = zeros(length(dfs)+1,length(Pvalues));
for i=1:length(dfs)
    for j=1:length(Pvalues)
       chiSquareMatrix(i+1,j) = chi2inv(1-Pvalues(j),dfs(i));
    end
end

fileID = fopen('chi-square-table.go','w');
fprintf(fileID,'package main\n\ntype chiSquareLookupTable struct {\n');
fprintf(fileID,'	P []float64\n');
fprintf(fileID,'	X [][]float64\n');
fprintf(fileID,'}\n\n');
fprintf(fileID,'var (\n// ChiSquareTable stores the frequently needed values for looking up Chi-square tests\nChiSquareTable = chiSquareLookupTable{\n')
fprintf(fileID,'[]float64{')
for j=1:length(Pvalues)
    if j==length(Pvalues)
        fprintf(fileID,'%2.6f},\n',Pvalues(j));
    else
        fprintf(fileID,'%2.6f, ',Pvalues(j));
    end
end
fprintf(fileID,'[][]float64{\n')
for i=1:length(dfs)
    fprintf(fileID,'[]float64{')
    for j=1:length(Pvalues)
        if j==length(Pvalues)
            fprintf(fileID,'%2.6f}, \n',chiSquareMatrix(i+1,j));
        else
            fprintf(fileID,'%2.6f, ',chiSquareMatrix(i+1,j))
        end
    end
end
fprintf(fileID,'		},\n	}\n)\n')
fclose(fileID);



Pvalues = [0.9999 0.999 0.995	0.975	0.20	0.10	0.05	0.025	0.02	0.01	0.005	0.002	0.001 0.0001];
Pvalues = sort(Pvalues);
Pvalues = unique(Pvalues);

fileID = fopen('normal-zero-one-table.go','w');
fprintf(fileID,'package main\n\ntype normalLookupTable struct {\n');
fprintf(fileID,'	P []float64\n');
fprintf(fileID,'	X []float64\n');
fprintf(fileID,'}\n\n');
fprintf(fileID,'var (\n// NormalTable stores the frequently needed values for looking up N(0,1) tests\nNormalTable = normalLookupTable{\n')
fprintf(fileID,'[]float64{')
for j=1:length(Pvalues)
    if j==length(Pvalues)
        fprintf(fileID,'%2.6f},\n',Pvalues(j));
    else
        fprintf(fileID,'%2.6f, ',Pvalues(j));
    end
end

fprintf(fileID,'[]float64{')
for j=1:length(Pvalues)
    if j==length(Pvalues)
        fprintf(fileID,'%2.6f}, \n',norminv(Pvalues(j),0,1));
    else
        fprintf(fileID,'%2.6f, ',norminv(Pvalues(j),0,1))
    end
end

fprintf(fileID,'		}\n)\n')
fclose(fileID);
