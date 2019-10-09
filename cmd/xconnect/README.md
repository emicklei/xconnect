# xconnect - command line tool

## extract to file

    xconnect -input some-configmap-application.properties.yml -k8s -target file://sample.yml

## POST to an HTTP endpoint

    xconnect -input some-configmap-application.properties.yml -k8s -target https://some-xconnect-handling-service.net
