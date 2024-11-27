// app/api/cards/interpret/[requestID]/route.ts
import { NextResponse } from "next/server";

export async function GET(
  request: Request,
  { params }: { params: { requestID: string } }
) {
  try {
    const response = await fetch(
      `http://localhost:8082/cards/interpret/${params.requestID}`
    );

    if (!response.ok) {
      throw new Error("Failed to retrieve tarot interpretation.");
    }

    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    console.error("Error retrieving interpretation:", error);
    return NextResponse.json(
      { error: "Error retrieving interpretation" },
      { status: 500 }
    );
  }
}
