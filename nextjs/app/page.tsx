import Image from "next/image";
import Counter from "./counter";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="flex flex-col items-center justify-center">
        <label className="text-4xl font-bold text-center"> Next.js </label>
        <Counter />
      </div>
    </main>
  );
}
