/**
 * 获取浏览器本地存储中的数据
 * @param {string} key - 要获取的存储项的键名
 * @returns {string | null} - 返回存储的数据，若没有则返回 null
 */
function GetFromLocalStorage(key: string) {
  const value = localStorage.getItem(key);
  return value ? value : null; // 若没有对应的值，则返回 null
}

/**
 * 设置浏览器本地存储中的数据
 * @param {string} key - 要设置的存储项的键名
 * @param {string} value - 要设置的值
 */
function SetToLocalStorage(key: string, value: string) {
  localStorage.setItem(key, value);
}

/**
 * 获取浏览器本地存储中的 token
 * @returns {string | null} - 返回 token，若没有则返回 null
 */
function GetToken() {
  const value = localStorage.getItem("token");
  return value ? value : null;
}

function RemoveToken() {
  localStorage.removeItem("token");
}

export { GetFromLocalStorage, SetToLocalStorage, GetToken, RemoveToken };
