cd ingress/deploy
echo "undeploy ingress"
./undeploy.ps1

cd ../../processor-analytics/deploy
echo "`nundeploy processor-analytics"
./undeploy.ps1

cd ../../service-transactional
echo "`nundeploy service-transactional"
./undeploy.ps1

cd ../service-analytics
echo "`nundeploy service-analytics"
./undeploy.ps1

cd ../processor-batch
echo "`nundeploy processor-batch"
./undeploy.ps1

cd ..