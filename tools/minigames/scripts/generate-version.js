const fs = require('fs');
const path = require('path');

// 生成版本号
const generateVersion = () => {
  const date = new Date();
  return `${date.getFullYear()}.${date.getMonth() + 1}.${date.getDate()}.${date.getHours()}${date.getMinutes()}`;
};

// 将版本号写入文件
const version = generateVersion();
const versionFile = path.join(__dirname, '..', 'src', 'version.js');

fs.writeFileSync(versionFile, `export const VERSION = '${version}';\n`);

console.log(`Version ${version} generated and saved to ${versionFile}`);
