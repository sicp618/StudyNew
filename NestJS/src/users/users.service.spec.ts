import { HttpException, HttpStatus } from '@nestjs/common';
import { UsersService } from './users.service';
import { User } from './user.schema';
import { getModelToken } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';

describe('UsersService', () => {
  let service: UsersService;
  let mockUserModel: any;

  const user = {
    username: 'test',
    password: '123456',
    email: 'test@test',
    id: 1,
  };

  const userDoc = {
    save: jest.fn().mockResolvedValue(user),
  };

  beforeEach(async () => {
    mockUserModel = {
      findOne: jest.fn(),
      create: jest.fn().mockReturnValue(userDoc),
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UsersService,
        {
          provide: getModelToken(User.name),
          useValue: mockUserModel,
        },
      ],
    }).compile();

    service = module.get<UsersService>(UsersService);
  });

  it('should register a new user', async () => {
    mockUserModel.create.mockReturnValueOnce(userDoc);

    expect(await service.register(user)).toEqual(user);
  });

  it('should throw an exception if the user already exists', async () => {
    mockUserModel.findOne.mockResolvedValue(user);

    await expect(service.register(user)).rejects.toThrow(new HttpException('User already exists', HttpStatus.CONFLICT));
    expect(mockUserModel.findOne).toHaveBeenCalledWith({ username: user.username });
  });
});
