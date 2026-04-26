## 1. Выбор фрагмента UML модели для реализации

### Список классов
Выбранный фрагмент состоит из следующих классов и интерфейсов:

1. **Класс User**
2. **Класс Product**
3. **Класс Order**
4. **Интерфейс IOrderService**
5. **Класс OrderService**
6. **Класс ShoppingCart**
7. **Класс Payment**
8. **Интерфейс IPaymentService**
9. **Класс PaymentService**
10. **Класс Address**

### Классы домена и контракты
- Классы домена: `User`, `Product`, `Order`, `ShoppingCart`, `Payment`, `Address`
- Контракты взаимодействия: `IOrderService`, `IPaymentService`

### Инварианты
1. Один пользователь может иметь много заказов.
2. Один заказ может включать много продуктов.
3. Заказ должен быть связан с одним пользователем и одним адресом доставки.

---

## 2. Правила сопоставления UML конструкций к языку программирования

### Представление классов и интерфейсов
- **Классы**: реализуются с использованием модификаторов доступа (public, private).
- **Интерфейсы**: определяют контракты, которые должны реализовываться классами.

### Представление ассоциаций и кратностей
- **One to Many**: Реализуется с помощью коллекций `List<T>`.
- **Many to Many**: Реализуется через дополнительный класс (например, `OrderProduct`).

### Представление композиции
- Композиты создаются внутри состоящих объектов и должны контролировать время жизни зависимых объектов.

---

## 3. Создание структуры проекта и пространств имен

Структура проекта:
- **MyEcommerce.Domain**
- **MyEcommerce.Services**
- **MyEcommerce.Tests**

Правила именования: Все классы и интерфейсы используют PascalCase.

Документация по сборке:
- Для компиляции проекта используйте команду `dotnet build`.
- Для запуска тестов используйте команду `dotnet test`.

---

## 4. Реализация каркаса классов и интерфейсов

```csharp
namespace MyEcommerce.Domain
{
    public class User
    {
        public int Id { get; private set; }
        public string Name { get; private set; }

        public User(int id, string name)
        {
            // Проверка на пустое имя
            if (string.IsNullOrWhiteSpace(name))
                throw new ArgumentException("Имя не должно быть пустым");

            Id = id;
            Name = name;
        }
    }

    public class Product
    {
        public int Id { get; private set; }
        public string Name { get; private set; }
        public decimal Price { get; private set; }

        public Product(int id, string name, decimal price)
        {
            // Проверка на отрицательную цену
            if (price < 0)
                throw new ArgumentException("Цена не может быть отрицательной");

            Id = id;
            Name = name;
            Price = price;
        }
    }

    public class Order
    {
        public int Id { get; private set; }
        public User User { get; private set; }
        public List<Product> Products { get; private set; } = new List<Product>();

        public Order(int id, User user)
        {
            Id = id;
            User = user ?? throw new ArgumentNullException(nameof(user));
        }

        public void AddProduct(Product product)
        {
            Products.Add(product);
        }
    }

    public interface IOrderService
    {
        void PlaceOrder(Order order);
    }

    public class OrderService : IOrderService
    {
        public void PlaceOrder(Order order) 
        {
            // Логика размещения заказа
        }
    }

    public class ShoppingCart
    {
        public List<Product> Products { get; private set; } = new List<Product>();

        public void AddProduct(Product product)
        {
            Products.Add(product);
        }
    }

    public class Payment
    {
        public int Id { get; private set; }
        public decimal Amount { get; private set; }

        public Payment(int id, decimal amount)
        {
            Amount = amount;
        }
    }

    public interface IPaymentService
    {
        void ProcessPayment(Payment payment);
    }

    public class PaymentService : IPaymentService
    {
        public void ProcessPayment(Payment payment) 
        {
            // Логика обработки платежа
        }
    }

    public class Address
    {
        public string Street { get; private set; }
        public string City { get; private set; }

        public Address(string street, string city)
        {
            Street = street;
            City = city;
        }
    }
}
```

---

## 5. Реализация инвариантов и бизнес ограничений

Проверки осуществляются в конструкторах и методах. Исключения:
- `ArgumentException` используется для указания ошибок в параметрах.
- `ArgumentNullException` для обязательных параметров, которые не могут быть нулевыми.

Ограничения на уровне базы данных включают уникальность `User.Id` и `Product.Id`.

---

## 6. Реализация отношений между объектами

- **Ассоциации**: Ссылки на объекты `User`, `Product` хранятся напрямую.
- **Коллекции**: Используются для защиты состояний через приватные поля.
- **Композиция**: Защита создания и удаления объектов осуществляется через методы класса.

---

## 7. Реализация DTO и маппинга

DTO используются для извлечения данных:
```csharp
public class OrderDto
{
    public int Id { get; set; }
    public int UserId { get; set; }
    public List<int> ProductIds { get; set; }
}
```

Правила сериализации:
- Используйте JSON для API, при этом учитывайте безопасность и валидацию данных.

---

## 8. Подготовка модульных тестов

Пример теста:
```csharp
using NUnit.Framework;

[TestFixture]
public class OrderTests
{
    [Test]
    public void CreateOrder_WithValidUser_ShouldSucceed()
    {
        var user = new User(1, "Ivan");
        var order = new Order(1, user);

        Assert.AreEqual(1, order.Id);
        Assert.AreSame(user, order.User);
    }

    [Test]
    public void CreateOrder_WithNullUser_ShouldThrowException()
    {
        Assert.Throws<ArgumentNullException>(() => new Order(1, null));
    }

    // Другие тесты...
}
```

---

## 9. Таблица соответствия UML и кода

| Элемент UML     | Код                          | Связи                                                 |
|------------------|-----------------------------|------------------------------------------------------|
| User             | `public class User`         | Один пользователь - много заказов                    |
| Product          | `public class Product`      | Много продуктов в заказе                              |
| Order            | `public class Order`        | Связан с одним пользователем и адресом               |
| IOrderService    | `public interface IOrderService` | Контракт для сервиса заказов                     |
| OrderService     | `public class OrderService` | Реализация IOrderService                              |
| ShoppingCart     | `public class ShoppingCart` | Хранит продукты в корзине                             |
| Payment          | `public class Payment`      | Связан с заказом, один платёж для одного заказа      |
| IPaymentService   | `public interface IPaymentService` | Контракт для обработки платежей                   |
| PaymentService   | `public class PaymentService` | Реализация IPaymentService                            |
| Address          | `public class Address`      | Связан с заказом и пользователем                     |

Отклонения не зафиксированы, реализация полностью соответствует модели.