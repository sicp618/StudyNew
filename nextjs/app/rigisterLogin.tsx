"use client";

import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { toast } from "react-toastify";

const registerUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/users/register`;
const loginUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/users/login`;

type User = {
  username: string;
  password: string;
  email: string;
};

export default function Counter() {
  const [user, setUser] = useState<User>({
    username: "",
    password: "",
    email: "",
  });

  const router = useRouter();

  const registerAction = async () => {
    const data = {
      user: {
        username: "user1",
        password: "123456",
        email: "user2@mail",
        id: 0,
      },
    };
    axios
      .post(registerUrl, data)
      .then((res) => {
        setUser(res.data.user);
      })
      .catch((error) => {
        console.log("error", error);
      });
  };

  const loginAction = async () => {
    const loginData = {
      user: {
        username: user.username,
        password: user.password,
        email: user.email,
        id: 0,
      },
    };
    axios
      .post(loginUrl, loginData)
      .then((res) => {
        console.log("res", res);
        setUser(res.data.user);

        toast.success("注册成功");
        router.push("/home");
      })
      .catch((error) => {
        toast.error(`注册失败, ${error}`);
      });
  };

  const actionSwitch = () => {
    console.log("actionSwitch");
  };

  return (
    <div>
      <div className="w-64">
        <input
          type="text"
          placeholder="username"
          value={user.username}
          onChange={(e) => setUser({ ...user, username: e.target.value })}
          className="input input-bordered w-full max-w-xs my-1"
        />
        <input
          type="password"
          placeholder="password"
          value={user.password}
          onChange={(e) => setUser({ ...user, password: e.target.value })}
          className="input input-bordered w-full max-w-xs my-1"
        />
        <button className="btn btn-primary w-full my-1" onClick={loginAction}>
          Login
        </button>
        <div className="form-control my-1">
          <label className="hover" onClick={registerAction}>
            register
          </label>
        </div>
      </div>
    </div>
  );
}
