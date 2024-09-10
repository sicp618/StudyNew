import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { PrismaService } from '../prisma/prisma.service';
import { User, Prisma } from '@prisma/client';

@Injectable()
export class UsersService {
  constructor(private prisma: PrismaService) {}

  async register(user: User): Promise<User> {
    const userExists = await this.getUserByUsername(user.username);
    if (userExists) {
      throw new HttpException('User already exists', HttpStatus.CONFLICT);
    }

    const emailExists = await this.getUserByEmail(user.email);
    if (emailExists) {
      throw new HttpException('Email already exists', HttpStatus.CONFLICT);
    }

    const createdUser = await this.createUser(user);
    return createdUser;
  }

  async login(user: User): Promise<User> {
    const userExists = await this.getUserByUsername(user.username);
    if (!userExists) {
      throw new HttpException('User does not exist', HttpStatus.NOT_FOUND);
    }

    if (userExists.password !== user.password) {
      throw new HttpException('Password is incorrect', HttpStatus.UNAUTHORIZED);
    }

    return userExists;
  }

  async createUser(data: Prisma.UserCreateInput): Promise<User> {
    return this.prisma.user.create({
      data,
    });
  }

  async getUser(id: number): Promise<User | null> {
    return this.prisma.user.findUnique({
      where: { id },
    });
  }

  async getUserByUsername(username: string): Promise<User | null> {
    return this.prisma.user.findUnique({
      where: { username },
    });
  }

  async getUserByEmail(email: string): Promise<User | null> {
    return this.prisma.user.findUnique({
      where: { email },
    });
  }

  async getUsers(): Promise<User[]> {
    return this.prisma.user.findMany({
      where: { deleted: false },
    });
  }

  async updateUser(id: number, data: Prisma.UserUpdateInput): Promise<User> {
    return this.prisma.user.update({
      where: { id },
      data,
    });
  }

  async deleteUser(id: number): Promise<User> {
    return this.prisma.user.update({
      where: { id },
      data: { deleted: true },
    });
  }
}
