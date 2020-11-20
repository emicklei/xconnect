# xconnect - command line tool

## validate a file

    xconnect -input some-configmap-application.properties.yaml

## extract to file

    xconnect -input some-configmap-application.properties.yaml -k8s -target file://sample.yaml

## POST to an HTTP endpoint

    xconnect -input some-configmap-application.properties.yaml -k8s -target https://some-xconnect-handling-service.net
