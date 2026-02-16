## Twitter LLD – Interview Walkthrough

This repository models a simplified Twitter using clean layering, SOLID principles and a couple of core design patterns.  
This README is written as if you are explaining your approach in a system design / LLD interview.

---

## 1. Problem Restatement (How to start in an interview)

**Start by clarifying requirements:**
- **Core entities**: `User`, `Tweet`, `Comment`.
- **Basic operations**:
  - Users can **tweet**.
  - Users can **like** and **comment** on tweets.
  - Users can **follow / unfollow** other users.
  - Users can **see a timeline/feed** of tweets.
- **Feed behaviors**:
  - Support **different feed strategies**, e.g. time-sorted, popularity-sorted.
  - The system should be extensible to add more feed strategies later.

**Then propose a layered structure:**
- **domain**: pure business entities (`User`, `Tweet`, `Comment`).
- **interfaces**: abstractions for services (`TweetService`, `UserService`, `FeedService`, `ILoadFeedStrategy`).
- **services**: concrete implementations of those interfaces.
- **usecase**: algorithm-level strategies for building feeds.
- **facade**: a `Twitter` object that exposes a simple API to the outside world.

This gives the interviewer confidence that you structure code with clear responsibilities before diving into details.

---

## 2. High-level Design & Responsibilities

- **Domain layer (`internal/domain`)**
  - `Tweet`, `User`, `Comment` are pure data + behavior:
    - `Tweet`: content, owner, time, likes, comments.
    - `User`: id, name, followers and operations to add/remove followers.
    - `Comment`: metadata about a comment.

- **Service layer (`internal/services`)**
  - `TweetService`:
    - Store tweets in memory.
    - `AddTweet`, `Like`, `AddComment`, `GetTweetMap`.
  - `UserService`:
    - Maintain users and follower relationships.
    - `AddUser`, `Follow`, `UnFollow`, `GetAllFollowers`.
  - `FeedService`:
    - Depends on `TweetService` **via interface** to get data.
    - Chooses feed strategy and loads timeline.

- **Use case layer (`internal/usecase`)**
  - `TimeSortedFeed` and `PopularitySortedFeed`:
    - Both implement `ILoadFeedStrategy`.
    - They encapsulate the algorithm for ordering tweets.

- **Facade (`internal/services/twitter.go`)**
  - `Twitter`:
    - A single entry point that exposes:
      - `Tweet`, `Comment`, `Like`
      - `Follow`, `UnFollow`
      - `LoadTimeline`
    - Internally holds:
      - `FeedService` (interface)
      - `TweetService` (interface)
      - `UserService` (interface)
    - Constructed as a **singleton** to simulate a single application-level instance.

---

## 3. SOLID Principles in this Codebase

### 3.1 Single Responsibility Principle (SRP)

- **`TweetService`**:
  - Focused only on tweet-related operations (create, like, comment, expose tweet map).
- **`UserService`**:
  - Focused only on user and follower management.
- **`FeedService`**:
  - Focused on building/serving feeds using strategies; it does not care how tweets are stored.
- **`TimeSortedFeed` / `PopularitySortedFeed`**:
  - Each class has a single responsibility: one way of ordering tweets.

In an interview, you can say:
- **“I’m keeping each service focused on one axis of behavior: tweet lifecycle, user graph, and feed generation.”**

### 3.2 Open/Closed Principle (OCP)

- **Feed strategies** are **open for extension, closed for modification**:
  - Existing code in `FeedService.GetFeedStrategy` returns an `ILoadFeedStrategy`.
  - To add a new strategy (e.g., “following-only feed” or “trending feed”):
    - You implement another struct that satisfies `ILoadFeedStrategy`.
    - Optionally extend the `FeedStrategy` enum and `GetFeedStrategy` switch.
  - The `LoadTimeline` function still just depends on `ILoadFeedStrategy` and doesn’t change.

In an interview:
- **“I designed feed generation using a strategy interface so I can add more algorithms without touching consumers.”**

### 3.3 Liskov Substitution Principle (LSP)

- **`ILoadFeedStrategy` implementations**:
  - `TimeSortedFeed` and `PopularitySortedFeed` both implement `LoadFeed() []*domain.Tweet`.
  - Anywhere the code expects `ILoadFeedStrategy`, either implementation (or any new one) can be substituted without breaking behavior.
- **Service interfaces (`TweetService`, `UserService`, `FeedService`)**:
  - The `Twitter` facade and `FeedService` structs depend only on these interfaces.
  - Any new implementation (e.g., **DB-backed tweet service**, **cached user service**) can replace the current in-memory ones.

How to phrase it:
- **“All my concrete services and strategies can be swapped as long as they respect the interface contracts, satisfying Liskov substitution.”**

### 3.4 Interface Segregation Principle (ISP)

- `TweetService` interface only exposes methods that consumers need:
  - Initially: `GetTweetMap()`.
  - For the facade: `AddTweet`, `Like`, `AddComment`.
- `ILoadFeedStrategy` exposes only `LoadFeed()`; it doesn’t leak any other behavior.

Explain in interview:
- **“Clients don’t depend on methods they don’t use. For example, feed logic only knows about `LoadFeed`, and services only expose exactly what their consumers need.”**

### 3.5 Dependency Inversion Principle (DIP)

- High-level modules depend on **interfaces**, not concrete types:
  - `FeedService` depends on `interfaces.TweetService`.
  - `Twitter` depends on `interfaces.FeedService`, `interfaces.TweetService`, `interfaces.UserService`.
- Concrete implementations (`services.TweetService`, `services.UserService`, `services.FeedService`) are created in composition root (`NewTwiter`) and injected behind those interfaces.

Interview phrasing:
- **“I inverted dependencies so the façade and feed logic depend on abstractions, enabling easier testing and swapping implementations.”**

---

## 4. Design Patterns Used

### 4.1 Strategy Pattern (Feed strategies)

- **Context**: `FeedService`.
- **Strategy interface**: `ILoadFeedStrategy` (`LoadFeed() []*Tweet`).
- **Concrete strategies**:
  - `TimeSortedFeed`: sorts tweets by time.
  - `PopularitySortedFeed`: sorts tweets by number of likes.

Why this is Strategy:
- The algorithm for building the feed is encapsulated in separate types and chosen at runtime based on `FeedStrategy` enum.

### 4.2 Facade Pattern (Twitter API)

- **`Twitter` struct** in `services`:
  - Hides complexity of multiple services (`TweetService`, `UserService`, `FeedService`).
  - Offers simple methods like `Tweet`, `Like`, `Comment`, `Follow`, `LoadTimeline`.

How to explain:
- **“I created a Twitter facade to give the client a simple API while hiding the internal service orchestration.”**

### 4.3 Singleton Pattern (Twitter instance)

- `NewTwiter` uses `sync.Once` and a package-level `TwitterInst` to ensure a single instance:
  - This is a classical Go singleton implementation.

Note: In production you’d be more careful with singletons, but for an LLD question it demonstrates that you know the pattern and how to implement it safely.

---

## 5. Step-by-step Interview Approach

Use this **sequence** when the interviewer asks you to design a “mini Twitter” with feeds.

### Step 1: Clarify Requirements and Scope

- Ask:
  - Do we need **persistence** or is in-memory OK for now?
  - Which operations must we support: tweeting, liking, commenting, following, feed?
  - How many types of feed sorting do we need now? Is it likely to change?
- Propose:
  - For an LLD round, keep storage in-memory and focus on **clean design and extensibility**.

### Step 2: Identify Core Entities (Domain)

- Sketch `User`, `Tweet`, `Comment`:
  - What fields do they have?
  - What behavior belongs inside them? (e.g., `Tweet.Like`, `Tweet.AddComment`, `User.AddFollower`).
- Mention SRP:
  - **“These domain entities encapsulate the core data and minimal behavior; higher-level coordination will live in services.”**

### Step 3: Define Services and Their Responsibilities

- Propose three services:
  - `TweetService`: manage tweets and operations on them.
  - `UserService`: manage users and followers.
  - `FeedService`: responsible for building and displaying timelines.
- Call out SRP and separation of concerns.

### Step 4: Introduce Abstractions (Interfaces) – Apply DIP/ISP

- Create interfaces in an `interfaces` package:
  - `TweetService`
  - `UserService`
  - `FeedService`
  - `ILoadFeedStrategy`
- Explain:
  - **“High-level modules (like the Twitter façade and feed logic) will depend on these interfaces instead of concrete implementations, to satisfy Dependency Inversion and make the design testable and extensible.”**

### Step 5: Design Feed Extensibility with Strategy Pattern

- Introduce `ILoadFeedStrategy` with method `LoadFeed() []*Tweet`.
- Define concrete strategies:
  - `TimeSortedFeed`.
  - `PopularitySortedFeed`.
- In `FeedService`, add `GetFeedStrategy(feedStrategy FeedStrategy) ILoadFeedStrategy` and `LoadTimeline`.
- Explain:
  - **“To support multiple feed algorithms and new ones in the future, I use the Strategy pattern via an `ILoadFeedStrategy` interface. Adding a new algorithm doesn’t change the consumers.”**

### Step 6: Build the Twitter Facade

- Design a `Twitter` struct that exposes the main API:
  - `Tweet`, `Like`, `Comment`, `Follow`, `UnFollow`, `LoadTimeline`.
- Internally it holds:
  - `FeedService interfaces.FeedService`
  - `TweetService interfaces.TweetService`
  - `UserService interfaces.UserService`
- Explain:
  - **“The façade pattern gives the client a simple unified interface; internally it orchestrates multiple services.”**

### Step 7: Wire Everything Together (Composition Root)

- In `NewTwiter`:
  - Instantiate concrete `TweetService`, `UserService`, `FeedService`.
  - Inject them into the `Twitter` facade as **interfaces**.
  - Use `sync.Once` to enforce singleton semantics.
- Highlight:
  - **DIP**: all dependencies go through interfaces.
  - **Singleton pattern**: one central `Twitter` instance for the demo.

### Step 8: Walk Through a Sample Flow

Using `cmd/main.go`, narrate:
- Create users `Alice` and `Bob`.
- Add them via `Twitter.UserService.AddUser`.
- `Follow`, `Tweet`, `Like`, `Comment`.
- Call `LoadTimeline` with a specific `FeedStrategy` (e.g., `TimeSortedFeed`).

Explain what happens behind the scenes:
- `Twitter` delegates to `TweetService` and `UserService`.
- `FeedService` decides which `ILoadFeedStrategy` to use.
- The chosen strategy orders tweets and returns them for printing.

This helps the interviewer see that you understand both the **static design** and the **runtime behavior**.

---

## 6. How to Talk About Trade-offs

If there’s time, mention:
- **In-memory storage**:
  - Good for LLD clarity and interview time.
  - Could later be swapped out for DB-backed implementations thanks to interfaces (DIP + LSP).
- **Thread safety / concurrency**:
  - Not fully handled here; in a production system, you’d add locks or use concurrent-safe structures.
- **Scalability**:
  - This design is focused on structure, not on massive-scale Twitter; but the use of clean boundaries makes it easier to evolve.

---

## 7. Running the Example

From the project root:

```bash
go run ./cmd
```

You should see:
- Users being created and following each other.
- Tweets being posted, liked, and commented on.
- A final timeline printed according to the selected feed strategy.

