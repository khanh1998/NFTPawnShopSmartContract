import { getRandomInt } from './math';

const colorNames = [
  'red', ' pink', ' purple', ' deep-purple', ' indigo', ' blue',
  ' light-blue', ' cyan', ' teal', ' green', ' light-green', ' lime',
  ' yellow', ' amber', ' orange', ' deep-orange', ' brown', ' blue-grey', ' grey',
];

export function getRandomColor(): string { // eslint-disable-line
  const rand = getRandomInt(0, colorNames.length);
  return colorNames[rand];
}
