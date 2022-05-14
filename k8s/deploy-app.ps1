cd ingress/deploy
echo "deploy ingress"
./deploy.ps1

cd ../../processor-analytics/deploy
echo "`ndeploy processor-analytics"
./deploy.ps1

cd ../../service-transactional
echo "`ndeploy service-transactional"
./deploy.ps1

cd ../service-analytics
echo "`ndeploy service-analytics"
./deploy.ps1

cd ../processor-batch
echo "`ndeploy processor-batch"
./deploy.ps1

cd ..