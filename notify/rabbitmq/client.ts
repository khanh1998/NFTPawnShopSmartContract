import amqp = require('amqplib/callback_api')
import { Server } from 'socket.io';

export class RabbitMQ {
	protected conn: any = null
	protected io: Server = null
	constructor(uri: string, channelName: string, io: Server) {
		this.io = io
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
						io.emit('data_update', JSON.parse(msg.content))
						}, { noAck: true }
					);
				});
			}
		});
	}


}