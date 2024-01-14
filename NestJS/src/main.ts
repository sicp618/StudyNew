import { NestFactory } from '@nestjs/core';
import type { NestExpressApplication } from '@nestjs/platform-express';
import { AppModule } from './app.module';

async function bootstrap() {
  try {
    const app = await NestFactory.create<NestExpressApplication>(AppModule, {
      cors: true,
    });
    await app.listen(process.env.PORT || 3000);
  } catch (error) {
    console.error('error during application start', error);
  }
}
bootstrap();
