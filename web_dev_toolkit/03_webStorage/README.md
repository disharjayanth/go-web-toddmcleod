use localStorage in client browser
localStorage.setItem("key", "value")
localStorage.getItem("key", "value")
localStorage.removeItem("key")

use sessionStorage for only that particular session. Once browser is closed session storage is lost.
sessionStorage.setItem("key", "value")
sessionStorage.getItem("key", "value")
sessionStorage.removeItem("key")
