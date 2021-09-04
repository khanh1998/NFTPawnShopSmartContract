import { NestFactory } from '@nestjs/core';
import { ConfigService } from '@nestjs/config';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableCors({
    origin: '*',
  });
  // const configService = app.get(ConfigService)
  // await app.connectMicroservice<MicroserviceOptions>({
  //   transport: Transport.GRPC,
  //   options: {
  //     package: 'app',
  //     protoPath: join(process.cwd(), 'src/app.proto'),
  //     url: configService.get('GRPC_CONNECTION_URL'),
  //   },
  // })
  // app.startAllMicroservices();
  await app.listen(5000);
}
bootstrap();
