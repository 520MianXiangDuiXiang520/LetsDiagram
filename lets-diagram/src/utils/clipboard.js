export async function copyCode(code) {
  try {
    await navigator.clipboard.writeText(code);
    console.log("Page URL copied to clipboard");
  } catch (err) {
    console.error("Failed to copy: ", err);
  }
}
