import ColorHash from "color-hash";

const colorHash = new ColorHash({ saturation: 1.0 });
const stringToColor = (str: string): string => colorHash.hex(str);

export const generateColor = (str: string): [string, string] => {
  const st1 = str.substring(0, str.length / 2);
  const st2 = str.substring(str.length / 2);
  const colorOne = stringToColor(st1);
  const colorTwo = stringToColor(st2);

  return [colorOne, colorTwo];
};

const generateImage = (str: string) => {
  const size = 512
  const [colorOne, colorTwo] = generateColor(str);
  const Image = `
    <svg width="${size}" height="${size}" viewBox="0 0 ${size} ${size}" fill="none" xmlns="http://www.w3.org/2000/svg">
      <circle cx="${size / 2}" cy="${size / 2}" r="${size / 2
    }" fill="url(#gradient)" />
      <defs>
        <linearGradient id="gradient" x1="0" y1="0" x2="${size}" y2="${size}" gradientUnits="userSpaceOnUse">
          <stop stop-color="${colorOne}" />
          <stop offset="1" stop-color="${colorTwo}" />
        </linearGradient>
      </defs>
    </svg>
      `.trim();

  return Image;
}

export default generateImage;