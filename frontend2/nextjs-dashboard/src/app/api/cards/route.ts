import axios from "axios";

export async function POST(request: Request) {
  try {
    const { name, query } = await request.json();
    const response = await axios.post("http://localhost:8082/api/cards", {
      name,
      query,
    });
    const { requestID } = response.data;
    return new Response(JSON.stringify({ requestID }), {
      status: 200,
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (error) {
    console.error("Error interpreting tarot cards:", error);
    return new Response(
      JSON.stringify({ error: "Failed to interpret tarot cards" }),
      {
        status: 500,
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  }
}
