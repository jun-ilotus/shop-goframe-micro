# GoFrame Template For MonoRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399
- https://goframe.org/pages/viewpage.action?pageId=87246764

docker run -d --name elasticsearch `
    -p 9200:9200 -p 9300:9300 `
    -e "discovery.type=single-node" `
    -e "xpack.security.enabled=false" `
    -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" `
    -v es-data:/usr/share/elasticsearch/data `
    docker.elastic.co/elasticsearch/elasticsearch:8.11.0

docker run -d --name kibana `
    -p 5601:5601 `
    -e "ELASTICSEARCH_HOST=http://elasticsearch:9200" `
    -e "SERVER_NAME=kibana" `
    -e "XPACK_SECURITY_ENABLED=false" `
    -e "XPACK_MONITORING_ENABLED=false" `
    --link elasticsearch:elasticsearch `
    docker.elastic.co/kibana/kibana:8.11.0