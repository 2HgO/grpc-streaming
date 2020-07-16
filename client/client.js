const {VideoInfo} = require('./pb/source_pb');
const {SourceClient} = require('./pb/source_grpc_web_pb');

const { UserInfo } = require('./pb/auth_pb');
const { AuthPromiseClient: AuthClient } = require('./pb/auth_grpc_web_pb');

const auth_client = new AuthClient("http://localhost:55099");
const video_client = new SourceClient("http://localhost:55099");

const __codec__ = 'video/mp4; codecs="avc1.640032"; profiles="mp42,mp41,isom,avc1"';
const EOF_INDICATOR = 0xffffffff

var token = "";

async function streamVideo() {
	const user_info = new UserInfo();
	user_info.setUsername("oghogho");
	user_info.setPassword("password");

	// Create user in case user does not exist
	await auth_client.createUser(user_info, {})
		.then(res => console.log(`User creation on auth-server status: ${res.getSuccess()}`))
		.catch((err) => console.log(`Error creating user: ${err.message}`));

	var video_info = new VideoInfo();
	video_info.setName("Daffodil");

	let err = null;
	const res = await auth_client.login(user_info, {}).catch(e => {
		err = e;
		return null;
	});

	if (err != null || res == null) {
		console.log("Error: ", err)
		return;
	}

	if (!res.getSuccess()) {
		console.log("User not authenticated");
		return;
	}
	token = res.getToken()

	var video = document.querySelector("video");
	var feed = new MediaSource;

	video.src = URL.createObjectURL(feed);
	var chunks = [];
	feed.addEventListener('sourceopen', function (_) {
		var buffer = feed.addSourceBuffer(__codec__)
		buffer.addEventListener('updateend', function() {
			if (chunks.length == 1 && chunks[0] == EOF_INDICATOR) {
				try { feed.endOfStream(); } catch (e) {}
				video.play();
				return;
			}
			if (!buffer.updating && chunks.length > 0) {
				buffer.appendBuffer(chunks.shift());
			}
		});
		console.log("started media source...");

		var stream = video_client.getVideo(video_info, { authorization:[token] });
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
}

streamVideo();
