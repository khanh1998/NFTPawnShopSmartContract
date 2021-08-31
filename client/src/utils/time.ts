export function convertSecondToDateStr(seconds: number): string {
  return new Date(seconds * 1000).toLocaleString();
}

export function convertSecondAndDurationToDateStr(seconds: number, duration: number): string {
  const date = new Date(seconds * 1000);
  date.setDate(date.getDate() + duration);
  return date.toLocaleString();
}
