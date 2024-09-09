import { NestFactory } from '@nestjs/core';
import type { NestExpressApplication } from '@nestjs/platform-express';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  try {
    const app = await NestFactory.create<NestExpressApplication>(AppModule, {
      cors: true,
    });
    const configService = app.get(ConfigService);
    await app.listen(configService.get<number>('PORT') || 4000);
  } catch (error) {
    console.error('error during application start', error);
  }
}
bootstrap();
