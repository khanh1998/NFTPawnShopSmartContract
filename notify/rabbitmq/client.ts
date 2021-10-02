import amqp = require('amqplib/callback_api')

export class RabbitMQ {
	protected conn: any = null
	constructor(uri: string, channelName: string, callback: (msg: any) => void) {
		amqp.connect(uri, function(error, connection) {
			if (error) {
				throw error
			} else {
				console.log('connect to rabbitmq successfully')
				connection.createChannel(function(error, channel) {
					if (error) {
						throw error;
					}
					channel.assertQueue(channelName, {
						durable: false
					});
					channel.consume(channelName, function(msg) {
						console.log(" [x] Received %s", msg.content.toString());
						}, { noAck: true }
					);
				});
			}
		});
	}


}