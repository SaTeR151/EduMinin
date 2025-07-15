// main.js
document.addEventListener('DOMContentLoaded', function() {
    // toggle-меню на мобильных
    const menuToggle = document.querySelector('.menu-toggle');
    const mainNav   = document.querySelector('.main-nav');

    menuToggle.addEventListener('click', () => {
    mainNav.classList.toggle('open');
});
    // Плавный скролл для якорных ссылок
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            e.preventDefault();
            document.querySelector(this.getAttribute('href'))
                    .scrollIntoView({ behavior: 'smooth' });
        });
    });

    // Изменение фона меню при скролле
    const header = document.querySelector('.site-header');
    window.addEventListener('scroll', function() {
        if (window.scrollY > 50) {
            header.classList.add('scrolled');
        } else {
            header.classList.remove('scrolled');
        }
    });
});

document.querySelectorAll('.news-container, .directions-container')
  .forEach(container => {
    container.addEventListener('wheel', e => {
      // если есть вертикальное движение колеса — прокручиваем влево/вправо
      if (Math.abs(e.deltaY) > Math.abs(e.deltaX)) {
        e.preventDefault();
        container.scrollLeft += e.deltaY;
      }
    });
});