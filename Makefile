.PHONY: start
start: video-server envoy-proxy static-gateway

.PHONY: video-server
video-server:
	@ cd services/videoStore && docker build -t stream/service .  &> /dev/null
	@ docker run --name video-server --rm -d --expose 55052 stream/service  &> /dev/null
	@ sleep 3

.PHONY: envoy-proxy
envoy-proxy:
	@ cd client && docker build -f Dockerfile.envoy -t stream/envoy .  &> /dev/null
	@ docker run --name envoy-proxy --rm -d -p 55099:55099 -p 9901:9901 --link video-server stream/envoy  &> /dev/null
	@ sleep 5

.PHONY: static-gateway
static-gateway:
	@ cd client && python3 -m http.server 8081

.PHONY: clean
clean:
	@ docker kill video-server | true &> /dev/null
	@ docker kill envoy-proxy | true &> /dev/null
	@ docker image rm stream/service | true &> /dev/null
	@ docker image rm stream/envoy | true &> /dev/null

transform-mp4:
	MP4Box -dash 1000 -rap -frag-rap ${FILE_NAME}