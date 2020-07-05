const {VideoInfo} = require('./source_pb');
const {SourceClient} = require('./source_grpc_web_pb');

const client = new SourceClient("http://localhost:55099");

var __codec__ = 'video/mp4; codecs="avc1.640032"; profiles="mp42,mp41,isom,avc1"';

var info = new VideoInfo();
info.setName("Dafodil");

const EOF_INDICATOR = 0xffffffff

// var writer = new WritableStream()

var video = document.querySelector("video");
var feed = new MediaSource;

video.src = URL.createObjectURL(feed);
var chunks = [];
feed.addEventListener('sourceopen', function (_) {
	var buffer = feed.addSourceBuffer(__codec__)
	buffer.addEventListener('updateend', function() {
		if (chunks.length == 1 && chunks[0] == EOF_INDICATOR) {
			feed.endOfStream();
			video.play();
			return;
		}
		if (!buffer.updating && chunks.length > 0) {
			buffer.appendBuffer(chunks.shift());
		}
	});
	console.log("started media source...");

	var stream = client.getVideo(info, {});
	stream.on('data', (res) => {
		if (buffer.updating || chunks.length) {
			chunks.push(res.getData());
			return;
		}
		buffer.appendBuffer(res.getData());
	});
	stream.on('error', (err) => {
		console.error(`Unexpected stream error: code = ${err.code}, message = ${err.message}`)
	});
	stream.on('end', (_) => {
		chunks.push(EOF_INDICATOR);
		console.log("Stream finished from server.")
	});
}, false);
