import { Test, TestingModule } from '@nestjs/testing';
import { UsersController } from './users.controller';
import { UsersService } from './users.service';
import { HttpException } from '@nestjs/common';
import { User } from './user.schema';
import { SessionService } from '../session/session.service';

describe('UsersController', () => {
  let usersController: UsersController;
  let usersService: UsersService;
  let sessionService: SessionService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [UsersController],
      providers: [UsersService, SessionService],
    })
      .overrideProvider(UsersService)
      .useValue({
        register: jest.fn(),
      })
      .overrideProvider(SessionService)
      .useValue({
        createSession: jest.fn(),
      })
      .compile();

    usersController = module.get<UsersController>(UsersController);
    usersService = module.get<UsersService>(UsersService);
    sessionService = module.get<SessionService>(SessionService);
  });

  it('should return registered user', async () => {
    const user: User = {
      username: 'registeredUser',
      password: 'password',
      email: 'registerUser1@test',
      id: 0,
    };
    const registeredUser: User = {
      username: 'registeredUser1',
      password: 'password',
      email: 'registerUser1@test',
      id: 1,
    };

    jest.spyOn(usersService, 'register').mockResolvedValue(registeredUser);

    expect(await usersController.register(user)).toEqual({
      user: registeredUser,
    });
  });

  it('should throw error when registration fails', async () => {
    const user: User = {
      username: 'registeredUser',
      password: 'password',
      email: 'registerUser1@test',
      id: 0,
    };
    const error = new HttpException('Error message', 500);

    jest.spyOn(usersService, 'register').mockRejectedValue(error);

    await expect(usersController.register(user)).rejects.toThrow(error);
  });
});
