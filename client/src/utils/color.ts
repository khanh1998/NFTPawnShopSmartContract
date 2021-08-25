const colorNames = ['red', ' pink', ' purple', ' deep-purple', ' indigo', ' blue', ' light-blue', ' cyan', ' teal', ' green', ' light-green', ' lime', ' yellow', ' amber', ' orange', ' deep-orange', ' brown', ' blue-grey', ' grey'];

/**
 * Returns a random integer between min (inclusive) and max (inclusive).
 * The value is no lower than min (or the next integer greater than min
 * if min isn't an integer) and no greater than max (or the next integer
 * lower than max if max isn't an integer).
 * Using Math.round() will give you a non-uniform distribution!
 */
function getRandomInt(minNum: number, maxNum: number): number {
  const min = Math.ceil(minNum);
  const max = Math.floor(maxNum);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

export function getRandomColor(): string { // eslint-disable-line
  const rand = getRandomInt(0, colorNames.length);
  return colorNames[rand];
}
