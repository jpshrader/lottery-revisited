# The Lottery Algorithm

While I was a software developer intern in college, I was given a task to develop a 'lottery' application for the company. The company would regularly give away extra tickets to sporting events, concerts, comedy shows, etc, but found that their giveaway method was too involved and wasn't a level playing field.

So me and my fellow interns get started on building this lottery application. We built out the website (ASP.NET), email notifications, a dashboard to register for events, a management page to administer the events and the drawings.

However, I took particular interest in the 'lottery' aspect of this problem and how the application's drawing logic works. I didn't want the drawings to be completely random but weighted towards people who have signed up for events but havent won them - conversly, I wanted it to be weighted against those who have previously won events. It's important to note that I was taking an Operating Systems class and were discussing process scheduling algorithms and how operating systems deal with things like [starvation](https://www.codingninjas.com/studio/library/starvation-in-os). Operating Systems use a [Multilevel Feedback Queue](https://en.wikipedia.org/wiki/Multilevel_feedback_queue#Process_scheduling) to manage and track the scheduling of processes.

Quick simplified TLDR on process scheduling and the Multilevel Feedback Queue: There are 64 priority queues (level) (0-63), each level of has their own time limit (which increases the further down the queue you go - ex. 0 has a lower timelimit than 1 and so on). All processes start at the top level (0) of the queue, but if the process exceeds the time limit without finishing, they task is bumped down to the next priority. If it finishes under the alloted time, it will get bumped up. When choosing the next process to run, the scheduling algoritm chooses the highest priority item available. This has the effect of prioritizing fast, responsive processes (usually an interactive, UI process) over slow, long-running processes (backround tasks, etc). Which makes sense, as the user of the computer doesn't want their UI lagging and getting blocked by a long running process like a file download or some batch process. However there are issues with these lower. This system isn't perfect, as sometimes this results in `starvation`, which operating systems use various techniques, like `aging`, to address this [but that's not important right now](https://youtu.be/VOmD-xqK2Es?si=wosAAymrYyseUWG5&t=8).

20 year old me thought:

> Hey, this is actually quite similar, maybe an 'inverted' multilevel feedback queue would work here.

So I decided on the following rules:

1. Users will be assigned a priority. All users default to a priority of 63.
2. If a user registers for a drawing but doesn't win, their priority will get increased by 1 (e.g. 63 => 62).
3. If a user wins a drawing, their priority will be reset to 63.
4. The drawing algorithm will start at the top of the queues (0), and work its way through to the bottom (63). If no winners are found, it restarts from the top.
5. Users are shuffled within each priority level.
6. The drawing likelihood increases the further down the levels you go, to help counter-balance the ordering bias.

Point #6 caused some confusion at the time:

> The lower the priority, the higher the likelihood you win? That doesn't make sense.

If you compared them side by side, yes. However, this doesn't factor in the ordering bias (advantage of going first)

```
Person A and Person B are both entered in the drawing.

Person A has an individual probability of 50%
Person B has an individual probability of 75%

However, because Person A goes first, the *actual* probabilities are:

P(A) = 1 * .5               = .5
P(B) = (1 - P(A)) * .75     = .375
```

## Revisited: Elegant or Complex?


