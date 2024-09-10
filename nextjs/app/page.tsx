'use client';

import Image from "next/image";
import Counter from "./rigisterLogin";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="flex flex-col items-center justify-center">
        <Main />
      </div>
    </main>
  );
}

function Main() {
  const router = useRouter();
  router.push("/login");

  return (
    <div>
      <h1>Welcome to the Main Component</h1>
    </div>
  );
}
