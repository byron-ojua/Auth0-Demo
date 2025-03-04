export function buildBaseUrl() {
  // If env variable is "PROD"
    if (process.env.ENV === "PROD") {
        return "";
    }

    return "http://localhost:8080/";
}