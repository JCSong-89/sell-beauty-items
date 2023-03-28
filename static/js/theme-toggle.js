const themeToggle = document.querySelector('.theme-toggle');

// Check if the user has a saved theme preference
if (getCookie('theme') === 'dark') {
    document.body.classList.add('dark');
} else {
    document.body.classList.add('light');
}

// Handle theme toggle button clicks
themeToggle.addEventListener('click', () => {
    if (document.body.classList.contains('dark')) {
        document.body.classList.remove('dark');
        document.body.classList.add('light');
        setCookie('theme', 'light', 365);
    } else {
        document.body.classList.remove('light');
        document.body.classList.add('dark');
        setCookie('theme', 'dark', 365);
    }
});

function getCookie(name) {
    const cookies = document.cookie.split('; ');
    for (const cookie of cookies) {
        const [cookieName, cookieValue] = cookie.split('=');
        if (cookieName === name) {
            return cookieValue;
        }
    }
    return null;
}

// Set the value of a cookie with a given name and expiration date
function setCookie(name, value, days) {
    const date = new Date();
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
    const expires = '; expires=' + date.toUTCString();
    document.cookie = name + '=' + value + expires + '; path=/';
}
