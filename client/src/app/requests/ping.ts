import axios from "axios";
import { buildBaseUrl } from "./base";

export const fetchPing = async (): Promise<string> => {
  // console.log("Base URL:", buildBaseUrl());

  // return "hello from server 2";

  try {


    const response = await axios.get(`${buildBaseUrl()}ping`);
    return response.data.message; // "hello from server 2"
  } catch (error) {
    console.error("Error fetching ping:", error);
    throw new Error("Failed to fetch ping");
  }
};
