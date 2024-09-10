import { UsersService } from './users.service';
import { Response } from 'express';
import { SessionService } from '../session/session.service';
import { Controller, Get, Post, Body, Param, Put, Delete, HttpException, HttpStatus, Res, Req } from '@nestjs/common';
import { User, Prisma } from '@prisma/client';

@Controller('api/users')
export class UsersController {
  constructor(
    private readonly usersService: UsersService,
    private readonly sessionService: SessionService
  ) {}

  @Post('register')
  async register(@Body('user') user: User): Promise<{ user: User }> {
    try {
      console.log('register input:', user);
      const registeredUser = await this.usersService.register(user);
      return { user: registeredUser };
    } catch (err) {
      if (err instanceof HttpException) {
        throw err;
      }
      throw new HttpException(err.message, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

  @Post('login')
  async login(
    @Res({ passthrough: true }) res: Response,
    @Body('user') user: User
  ): Promise<{ user: User }> {
    try {
      const loggedInUser = await this.usersService.login(user);

      let sessionId = await this.sessionService.createSession(user.username);
      res.cookie('nid', sessionId, {
        httpOnly: true,
        // secure: true,
        maxAge: 1000 * 60 * 60 * 24 * 7,
      });
      return { user: loggedInUser };
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }
      throw new HttpException(error.message, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

  @Post('logout')
  async logout(
    @Req() req,
    @Res({ passthrough: true }) res: Response
  ): Promise<void> {
    try {
      await this.sessionService.deleteSession(req.cookies.nid);
      res.clearCookie('nid');
      res.status(200).send();
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }
      throw new HttpException(error.message, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }
}