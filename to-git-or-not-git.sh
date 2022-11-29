returnValue=`curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"serignmbaye\"}}){id}}"}'`
echo $returnValue | grep -Po '8521'