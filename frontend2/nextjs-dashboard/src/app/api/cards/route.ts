// app/api/cards/route.ts
import { NextResponse } from "next/server";

export async function POST(request: Request) {
  const { name, userStory } = await request.json();

  try {
    const response = await fetch(
      `http://localhost:8082/cards?name=${encodeURIComponent(
        name
      )}&userstory=${encodeURIComponent(userStory)}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to connect to tarot reading service.");
    }

    const data = await response.json();
    return NextResponse.json({ requestID: data.requestID });
  } catch (error) {
    console.error("Error fetching tarot reading:", error);
    return NextResponse.json(
      { error: "Error fetching tarot reading" },
      { status: 500 }
    );
  }
}
