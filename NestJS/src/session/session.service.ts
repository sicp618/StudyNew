import { Injectable } from '@nestjs/common';
import { v4 as uuidv4 } from 'uuid';
import { RedisService } from '../redis/redis.service';

@Injectable()
export class SessionService {
  constructor(private readonly redisService: RedisService) {}

  async createSession(userId: string): Promise<string> {
    const sessionId = uuidv4();
    console.log('sessionId:', sessionId);
    this.redisService.set(sessionId, userId, 60 * 60 * 24 * 30);

    return sessionId;
  }

  async getSession(sessionId: string): Promise<string | null> {
    return await this.redisService.get(sessionId);
  }

  async deleteSession(sessionId: string): Promise<void> {
    await this.redisService.delete(sessionId);
  }
}
