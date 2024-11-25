// app/page.tsx
"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function HomePage() {
  const [name, setName] = useState("");
  const [userStory, setUserStory] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch("/api/cards", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, userStory }),
      });

      if (response.ok) {
        const data = await response.json();
        router.push(`/cards/${data.requestID}`);
      } else {
        console.error("Failed to fetch tarot cards.");
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h1>Welcome to the Tarot Reading</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Your Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <textarea
          placeholder="Your Story"
          value={userStory}
          onChange={(e) => setUserStory(e.target.value)}
        />
        <button type="submit">Get Your Reading</button>
      </form>
    </div>
  );
}
