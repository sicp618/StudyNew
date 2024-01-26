"use client";

import axios from "axios";
import Router from "next/router";
import { useState } from "react";

const registerUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/register`;
const loginUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/login`;

type User = {
  username: string;
  password: string;
  email: string;
  id: number;
};


export default function Counter() {
  const [user, setUser] = useState<User | null>(null);
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

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
        username: username,
        password: password,
        email: "",
        id: 0,
      },
    };
    axios
      .post(loginUrl, loginData)
      .then((res) => {
        console.log("res", res);
        setUser(res.data.user);

        Router.push("/dashboard");
      })
      .catch((error) => {
        console.log("error", error);
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
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          class="input input-bordered w-full max-w-xs my-1"
        />
        <input
          type="password"
          placeholder="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          class="input input-bordered w-full max-w-xs my-1"
        />
        <button class="btn btn-primary w-full my-1" onClick={loginAction}>
          Login
        </button>
        <div class="form-control my-1">
          <label class="hover" onClick={registerAction}>
            register
          </label>
        </div>
      </div>
    </div>
  );
}
