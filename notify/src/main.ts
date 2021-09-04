import { NestFactory } from '@nestjs/core';
import { ConfigService } from '@nestjs/config';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService)
  await app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.GRPC,
    options: {
      package: 'app',
      protoPath: join(process.cwd(), 'src/app.proto'),
      url: configService.get('GRPC_CONNECTION_URL'),
    },
  })
  await app.startAllMicroservices();
  // for socket
  const app1 = await NestFactory.create(AppModule);
  app1.enableCors();
  await app1.listen(7789);
}
bootstrap();
