<?php
session_start();

// Данные пользователя (пример)
$user = [
    'name'  => 'Иван Иванов',
    'email' => 'ivan.ivanov@example.com',
    'phone' => '+7 (912) 345-67-89',
    'photo' => 'assets/images/user3.jpg'
];

// Список курсов пользователя
$myCourses = [
    ['img'=>'assets/icons/python.png',    'title'=>'Python для начинающих'],
    ['img'=>'assets/icons/1c.png',        'title'=>'1С: Профессионал'],
    ['img'=>'assets/icons/cpp_csharp.png','title'=>'C++/C# разработка'],
    ['img'=>'assets/icons/java.png',      'title'=>'Java-разработка'],
];
?>
<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Личный кабинет — Учебный центр</title>
  <link rel="icon" href="assets/icons/favicon.png" type="image/png">
  <link rel="stylesheet" href="css/style.css">
  <style>
    /* Вертикальный список «Мои курсы» */
    .my-courses-list {
      display: flex;
      flex-direction: column;
      gap: 20px;
    }
    .course-item {
      display: flex;
      align-items: center;
      background: #fff;
      border-radius: var(--border-radius);
      box-shadow: 0 2px 5px rgba(0,0,0,0.1);
      overflow: hidden;
    }
    .course-item img {
      flex-shrink: 0;
      width: 60px;
      height: auto;
      object-fit: cover;
      margin: 20px;
    }
    .course-item .course-info {
      padding: 15px;
      flex: 1;
    }
    .course-item .course-info h4 {
      margin: 0 0 5px;
      color: var(--primary-2);
    }
    .course-item .course-info p {
      margin: 0;
      color: #555;
      font-size: 0.9rem;
    }
  </style>
</head>
<body>
  <!-- фиксированное меню -->
  <header class="site-header">
    <nav class="main-nav">
      <button class="menu-toggle" aria-label="Меню">
        <img src="assets/icons/menu.png" alt="Меню" class="menu-icon">
      </button>
      <ul class="nav-links">
          <li><a href="index.html" class="active">Главная</a></li>
          <li><a href="news.html">Новости</a></li>
          <li><a href="reviews.html">Отзывы</a></li>
          <li><a href="signup.html">Запись</a></li>
          <li><a href="about.html">О нас</a></li>
      </ul>
      <ul class="auth-links">
        <li>
          <a href="lk.php" class="active">
            <?= htmlspecialchars($user['name'], ENT_QUOTES) ?: 'Личный кабинет' ?>
          </a>
        </li>
      </ul>
    </nav>
  </header>

  <main style="display:flex; gap:20px; max-width:1200px; margin:120px auto 60px; padding:0 20px;">
    <!-- Сайдбар -->
    <aside class="profile-sidebar" style="
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      flex: 0 0 250px;
      background: #fff;
      border-radius: var(--border-radius);
      box-shadow: 0 2px 5px rgba(0,0,0,0.1);
      padding: 20px;
      height: 100%; /* или фиксированная высота, если нужно */
    ">
  <!-- Верхний блок: аватар и основные кнопки -->
  <div>
    <img src="<?= htmlspecialchars($user['photo'], ENT_QUOTES) ?>" 
           alt="Аватар <?= htmlspecialchars($user['name'], ENT_QUOTES) ?>" 
           style="width: 95%; height: auto; border-radius: var(--border-radius); object-fit: cover; margin-bottom: 15px; display: block; margin-left: auto; margin-right: auto;">
    <button style="
           display:block; width:100%; margin-bottom:10px;
           padding:10px; background:var(--secondaryA-4);
           color:#fff; border:none; border-radius:var(--border-radius);
           cursor:pointer;
         ">
      Сменить фото
    </button>
    <button style="
           display:block; width:100%; margin-bottom:10px;
           padding:10px; background:var(--secondaryB-4);
           color:#fff; border:none; border-radius:var(--border-radius);
           cursor:pointer;
         ">
      Сменить имя
    </button>
  </div>

  <!-- Нижний блок: кнопка выйти -->
  <form action="logout.php" method="post" style="margin-top:20px;">
    <button type="submit" style="
           display:block; width:100%;
           padding:10px; background:#e74c3c;
           color:#fff; border:none;
           border-radius:var(--border-radius);
           font-weight:600;
           cursor:pointer; transition:background .3s;
         ">
      Выйти из личного кабинета
    </button>
  </form>
</aside>

    <!-- Основной контент -->
    <section style="flex:1; background:#fff; border-radius:var(--border-radius); box-shadow:0 2px 5px rgba(0,0,0,0.1); padding:20px;">
      <h2 style="margin:0 0 10px; color:var(--primary-2);">
        <?= htmlspecialchars($user['name'], ENT_QUOTES) ?>
      </h2>
      <p style="margin:0 0 5px;"><strong>Email:</strong> <?= htmlspecialchars($user['email'], ENT_QUOTES) ?></p>
      <p style="margin:0 0 30px;"><strong>Телефон:</strong> <?= htmlspecialchars($user['phone'], ENT_QUOTES) ?></p>

      <h3 style="margin-bottom:15px; color:var(--primary-3);">Мои курсы</h3>
      <div class="my-courses-list">
        <?php foreach ($myCourses as $course): ?>
          <div class="course-item">
            <img src="<?= htmlspecialchars($course['img'], ENT_QUOTES) ?>"
                 alt="<?= htmlspecialchars($course['title'], ENT_QUOTES) ?>">
            <div class="course-info">
              <h4><?= htmlspecialchars($course['title'], ENT_QUOTES) ?></h4>
              <p>Описание курса или дополнительные детали…</p>
            </div>
          </div>
        <?php endforeach; ?>
      </div>
    </section>
  </main>

  <!-- Подвал -->
  <footer class="site-footer">
        <div class="footer-top">
            <div class="footer-logo-desc">
                <img src="assets/images/footer-logo.png" alt="Логотип" class="footer-logo">
                <p>CodeBrothers — учебный IT-центр из Нижнего Новгорода.</p>
                <p>Учим программировать вместе и по-братски.</p>
            </div>
            <div class="footer-links">
                <a href="catalog.html">Каталог курсов</a>
                <a href="schedule.html">Расписание</a>
                <a href="signup.html">Запись</a>
                <a href="about.html">О нас</a>
                <a href="reviews.html">Отзывы</a>
            </div>
            <div class="footer-links" id="centers">
                <h4>Учебные центры</h4>
                <a href="centers.html">ул. Ульянова, 1</a>
                <a href="centers.html">ул. Ульянова, 10а</a>
                <a href="centers.html">ул. Челюскинцев, 9</a>
            </div>
        </div>
        <hr>
        <div class="footer-bottom">
            <div class="footer-copy">© 2025 Все права защищены</div>
            <div class="footer-policy">
                <a href="docs/Соглашение на обработку персональных данных.pdf">Соглашение на обработку персональных данных</a>
                <a href="docs/Политика конфиденциальности.pdf">Политика конфиденциальности</a>
                <a href="docs/Сведения об образовательной организации.pdf">Сведения об образовательной организации</a>
                <a href="docs/Информация об оплате.pdf">Информация об оплате</a>
            </div>
            <div class="footer-social">
                <a href="https://vk.com" target="_blank"><i class="fab fa-vk"></i></a>
                <a href="https://t.me" target="_blank"><i class="fab fa-telegram"></i></a>
            </div>
        </div>
    </footer>

  <script src="js/main.js"></script>
</body>
</html>
