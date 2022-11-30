key=$(grep "Church" -r interviews -l | cut -d "-" -f2)
echo $key
cat interviews/interview-"$key"
echo $MAIN_SUSPECT