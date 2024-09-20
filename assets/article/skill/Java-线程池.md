
# Java 线程池  
文章封面:  https://veport.oss-cn-beijing.aliyuncs.com/articles/f9fa18da262910eb13f802b003147915.jpg   
文章类型: 1   
文章分类: 学习   
文章标签: [线程池]   
创建时间: 2022-03-11 02:59:54 +0800 CST   

文章内容:
### 1 ThreadPoolExecutor 的构造方法中各个参数的含义是什么？

```java
public ThreadPoolExecutor(int corePoolSize,						// 第 1 个参数
                          int maximumPoolSize,					// 第 2 个参数
                          long keepAliveTime,					// 第 3 个参数
                          TimeUnit unit,						// 第 4 个参数
                          BlockingQueue<Runnable> workQueue,	// 第 5 个参数
                          ThreadFactory threadFactory,			// 第 6 个参数
                          RejectedExecutionHandler handler) {   // 第 7 个参数

```

| 序号 |      参数名       |          参数类型          | 参数含义                           |        取值范围         | 解释说明                       |
| ---- | :---------------: | :------------------------: | ---------------------------------- | :---------------------: | :----------------------------------------------------------- |
| 1    |  `corePoolSize`   |           `int`            | 核心线程数                         |           >=0           | 如果等于 0，则任务执行完成后，没有任何请求进入时就销毁线程池的线程；如果大于 0，即使本地任务执行完毕，核心线程也不会销毁。这个值的设置非常关键，设置过大会浪费资源，设置过小会导致线程频繁地创建或销毁。 |
| 2    | `maximumPoolSize` |           `int`            | 线程池能够容纳同时执行的最大线程数 | >0并且>= `corePoolSize` | 如果任务数超过第 5 个参数 `workQueue` 的任务缓存上限且待执行的线程数小于 `maximumPoolSize` 时，需要借助第 5 个参数的帮助，缓存在队列中。如果 `maximumPoolSize` 与 `corePoolSize` 相等，则是固定大小线程池。**最大线程数 = 核心线程数 + 非核心线程数。** |
| 3    |  `keepAliveTime`  |           `long`           | 线程池中的线程空闲时间             |           >=0           | 当空闲时间达到 `keepAliveTime` 值时，非核心线程会被销毁，直到只剩下 `corePoolSize` 个线程为止，避免浪费内存和句柄资源。在默认情况下，当线程池的线程数大于 `corePoolSize` 时，`keepAliveTime` 才会起作用。但是当 `ThreadPoolExecutor` 的 `allowCoreThreadTimeOut` 变量设置为 `true` 时，核心线程超时后也会被回收。 |
| 4    |      `unit`       |         `TimeUnit`         | 时间单位                           |                         | `keepAliveTime` 的时间单位通常是 `TimeUnit.SECONDS`。        |
| 5    |    `workQueue`    | `BlockingQueue<Runnable>`  | 任务缓存队列                       |     不可以为 `null`     | 当请求的线程数大于等于 `corePoolSize` 时，任务会进入 `BlockingQueue` 阻塞队列等待执行。 |
| 6    |  `threadFactory`  |      `ThreadFactory`       | 线程工厂                           |     不可以为 `null`     | 用来生成一组相同任务的线程。线程池的命名是通过给这个 factory 增加组名前缀来实现的。在虚拟机栈分析时，就可以知道线程任务是由哪个线程工厂产生的。 |
| 7    |     `handler`     | `RejectedExecutionHandler` | 执行拒绝策略的对象                 |     不可以为 `null`     | 当待执行的线程数大于等于 `maximumPoolSize` 时，就可以通过该策略处理请求，这是一种简单的限流保护。 |

### 2 ThreadPoolExecutor 执行任务的规则是什么？

1. 如果线程池中的线程数量未达到核心线程的数量，那么会直接启动一个核心线程来执行任务；
2. 如果线程池中的线程数量已经达到或者超过核心线程的数量，那么任务会被插入到任务队列中排队等待执行；
3. 如果在步骤 2 中无法将任务插入到任务队列中（原因往往是任务队列已满），这个时候如果线程数量未达到线程池规定的最大值，那么会启动一个非核心线程来执行任务；
4. 如果步骤 3 中线程数量已经达到线程池规定的最大值，那么就拒绝执行此任务。

绘制流程图如下：

![在这里插入图片描述](https://img-blog.csdnimg.cn/710c04bdff7d4cc8b711d3c1b5964d01.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAd2lsbHdheXdhbmc2,size_16,color_FFFFFF,t_70,g_se,x_16)

### 3 Executors 里的线程池有哪些？这些线程池有什么特点？

| 比较项          | newCachedThreadPool            | newFixedThreadPool(int nThreads)                             | newSingleThreadExecutor        | newScheduledThreadPool(int corePoolSize)                     |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| corePoolSize    | `0`                            | `nThreads`                     | 1    | corePoolSize                   |
| maximumPoolSize | `Integer.MAX_VALUE`            | `nThreads`                     | 1    | Integer.MAX_VALUE              |
| keepAliveTime   | `60L`                          | `0L`                           | `0L`                           | 0    |
| unit            | `TimeUnit.SECONDS`             | `TimeUnit.MILLISECONDS`        | `TimeUnit.MILLISECONDS`        | TimeUnit.NANOSECONDS           |
| workQueue       | `new SynchronousQueue<Runnable>()`                           | `new LinkedBlockingQueue<Runnable>()`                        | `new LinkedBlockingQueue<Runnable>()`                        | new DelayedWorkQueue()         |
| 线程池名称      | 无界线程池，可以进行自动线程回收                             | 固定大小线程池                 | 单线程线程池                   | 执行定时任务，重复任务线程池       |
| 特点            | 没有核心线程，只有非核心线程（最大为`Integer.MAX_VALUE`），超过 60s 的空闲线程会被回收，SynchronousQueue 会将任务直接提交给线程而不保持它们，所以任务会立即执行。 | 只有固定个数的核心线程，没有非核心线程，核心线程不会被回收，任务队列大小没有限制。当线程处于空闲状态时，不会被回收；当所有的线程都处于活动状态时，新任务到达就处于等待状态，直到有线程空闲出来。 | 只有一个核心线程，没有非核心线程，核心线程不会被回收，任务队列大小没有限制。可以保证所有的任务都在同一个线程中按顺序执行。 | 核心线程数是固定的，非核心线程数是没有限制的，非核心线程空闲时会被立即回收。 |





### 4 最后

对线程池的简单理解：

　　假如有一个工厂，工厂里面有10个工人，每个工人同时只能做一件任务。

　　因此只要当10个工人中有工人是空闲的，来了任务就分配给空闲的工人做；

　　当10个工人都有任务在做时，如果还来了任务，就把任务进行排队等待；

　　如果说新任务数目增长的速度远远大于工人做任务的速度，那么此时工厂主管可能会想补救措施，比如重新招4个临时工人进来；

　　然后就将任务也分配给这4个临时工人做；

　　如果说着14个工人做任务的速度还是不够，此时工厂主管可能就要考虑不再接收新的任务或者抛弃前面的一些任务了。

　　当这14个工人当中有人空闲时，而新任务增长的速度又比较缓慢，工厂主管可能就考虑辞掉4个临时工了，只保持原来的10个工人，毕竟请额外的工人是要花钱的。

 

　　这个例子中的corePoolSize就是10，而maximumPoolSize就是14（10+4）。

　　也就是说corePoolSize就是线程池大小，maximumPoolSize在我看来是线程池的一种补救措施，即任务量突然过大时的一种补救措施。

### 参考

- [《Android开发艺术探索》第11章-Android 的线程和线程池](https://blog.csdn.net/willway_wang/article/details/122632665?spm=1001.2014.3001.5502)
- [ Java并发编程：线程池的使用](https://www.cnblogs.com/dolphin0520/p/3932921.html)



### 测试：
```Java
package threadpool;


import java.util.concurrent.*;

/**
 * @Description create for javaCourse .
 * 这里要重点解释一下corePoolSize、maximumPoolSize、largestPoolSize三个变量。
 *
 * 　　corePoolSize在很多地方被翻译成核心池大小，其实我的理解这个就是线程池的大小。举个简单的例子：
 * 　　假如有一个工厂，工厂里面有10个工人，每个工人同时只能做一件任务。
 *
 * 　　因此只要当10个工人中有工人是空闲的，来了任务就分配给空闲的工人做；
 *
 * 　　当10个工人都有任务在做时，如果还来了任务，就把任务进行排队等待；
 *
 * 　　如果说新任务数目增长的速度远远大于工人做任务的速度，那么此时工厂主管可能会想补救措施，比如重新招4个临时工人进来；
 *
 * 　　然后就将任务也分配给这4个临时工人做；
 *
 * 　　如果说着14个工人做任务的速度还是不够，此时工厂主管可能就要考虑不再接收新的任务或者抛弃前面的一些任务了。
 *
 * 　　当这14个工人当中有人空闲时，而新任务增长的速度又比较缓慢，工厂主管可能就考虑辞掉4个临时工了，只保持原来的10个工人，毕竟请额外的工人是要花钱的。
 *
 *
 *
 * 　　这个例子中的corePoolSize就是10，而maximumPoolSize就是14（10+4）。
 *
 * 　　也就是说corePoolSize就是线程池大小，maximumPoolSize在我看来是线程池的一种补救措施，即任务量突然过大时的一种补救措施。
 * @Author weiyi
 * @Date 2022/3/11
 */
public class ThreadPoolTest {
    public static void main(String[] args) {

        //60s不执行就删，不够再加，无上限
        //cacheThreadPool();
        //一直添加线程，直到保持在固定数量
       // fixTheadPoolTest();
        //核心线程只有一个，保证先进先出
        //singleTheadPoolTest();
        //固定大小线程池，定时任务：循环执行、延迟执行
        //scheduleThreadPool();

        newThreadPool();
    }

    public static void newThreadPool() {
        /**
         * ArrayBlockingQueue; 规定大小的BlockingQueue，其构造必须指定大小。其所含的对象是FIFO顺序排序的。
         * LinkedBlockingQueue; 大小不固定的BlockingQueue，若其构造时指定大小，生成的BlockingQueue有大小限制，不指定大小，其大小有Integer.MAX_VALUE来决定。其所含的对象是FIFO顺序排序的。
         * PriorityBlockingQueue：类似于LinkedBlockingQueue，但是其所含对象的排序不是FIFO，而是依据对象的自然顺序或者构造函数的Comparator决定。
         * SynchronousQueue; 特殊的BlockingQueue，对其的操作必须是放和取交替完成。
         */
        BlockingQueue workQueue=new SynchronousQueue<Runnable>();

        /**
         * ThreadPoolExecutor.AbortPolicy:丢弃任务并抛出RejectedExecutionException异常。
         * ThreadPoolExecutor.DiscardPolicy：也是丢弃任务，但是不抛出异常。
         * ThreadPoolExecutor.DiscardOldestPolicy：丢弃队列最前面的任务，然后重新尝试执行任务（重复此过程）
         * ThreadPoolExecutor.CallerRunsPolicy：由调用线程处理该任务
         */
        ThreadFactory threadFactory=new ThreadFactory() {
            @Override
            public Thread newThread(Runnable r) {
               System.out.println("自定义线程工厂创建"+r.toString());
                Thread t = new Thread(r);
                t.setDaemon(true);
                return t;
            }
        };

        RejectedExecutionHandler rejectedExecutionHandler=new ThreadPoolExecutor.CallerRunsPolicy();
        ThreadPoolExecutor threadPoolExecutor = new ThreadPoolExecutor(
                3,
                5,
                60L,
                TimeUnit.SECONDS,
                workQueue,
                threadFactory,
                rejectedExecutionHandler);

        for (int i = 0; i < 5; i++) {
            Runnable r1 = new Runnable() {
                @Override
                public void run() {
                   System.out.println("Runnable 线程名称：" + Thread.currentThread().getName() + "，执行:3秒后执行");
                }
            };

            Callable c1=new Callable() {
                @Override
                public Object call() throws Exception {
                    System.out.println("Callable 线程名称：" + Thread.currentThread().getName() + "，执行:3秒后执行");
                    return 1234567;
                }
            };
            //execute无返回值 ——实现Runnable接口
            threadPoolExecutor.execute(r1);
            //submit有返回值 ——实现Callable接口
           System.out.println("------->"+threadPoolExecutor.submit(c1).toString());
        }
    }
    /**
     * newCachedThreadPool：
     * <p>
     * 底层：返回ThreadPoolExecutor实例，corePoolSize为0；
     * maximumPoolSize为Integer.MAX_VALUE；keepAliveTime为60L；时间单位TimeUnit.SECONDS；
     * workQueue为SynchronousQueue(同步队列)
     *
     * 通俗：当有新任务到来，则插入到SynchronousQueue中，由于SynchronousQueue是同步队列，因此会在池中寻找可用线程来执行，若有可以线程则执行，若没有可用线程则创建一个线程来执行该任务；若池中线程空闲时间超过指定时间，则该线程会被销毁。
     * 适用：执行很多短期的异步任务
     * 1.创建一个可缓存的线程池。如果线程池的大小超过了处理任务所需要的线程，那么就会回收部分空闲（60秒不执行任务）的线程<br>
     * 2.当任务数增加时，此线程池又可以智能的添加新线程来处理任务<br>
     * 3.此线程池不会对线程池大小做限制，线程池大小完全依赖于操作系统（或者说JVM）能够创建的最大线程大小<br>
     */
    public static void cacheThreadPool() {
        ExecutorService cachedThreadPool = Executors.newCachedThreadPool();
        for (int i = 1; i <= 10; i++) {
            final int ii = i*10;
            try {
                Thread.sleep(ii);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            cachedThreadPool.execute(new Runnable() {
                @Override
                public void run() {
                   System.out.println("线程名称：" + Thread.currentThread().getName() + "，执行" + ii);
                }
            });
        }
    }

    /**
     * newFixedThreadPool：
     *
     * 底层：返回ThreadPoolExecutor实例，接收参数为所设定线程数量n，
     * corePoolSize和maximumPoolSize均为n；keepAliveTime为0L；
     * 时间单位TimeUnit.MILLISECONDS；WorkQueue为：new LinkedBlockingQueue<Runnable>() 无界阻塞队列
     *
     * 通俗：创建可容纳固定数量线程的池子，每个线程的存活时间是无限的，当池子满了就不再添加线程了；如果池中的所有线程均在繁忙状态，对于新任务会进入阻塞队列中(无界的阻塞队列)
     * 适用：执行长期任务
     * 1.创建固定大小的线程池。每次提交一个任务就创建一个线程，直到线程达到线程池的最大大小<br>
     * 2.线程池的大小一旦达到最大值就会保持不变，如果某个线程因为执行异常而结束，那么线程池会补充一个新线程<br>
     * 3.因为线程池大小为3，每个任务输出index后sleep 2秒，所以每两秒打印3个数字，和线程名称<br>
     */
    public static void fixTheadPoolTest() {
        ExecutorService fixedThreadPool = Executors.newFixedThreadPool(3);
        for (int i = 0; i < 10; i++) {
            final int ii = i;
            fixedThreadPool.execute(() -> {
               System.out.println("线程名称：" + Thread.currentThread().getName() + "，执行" + ii);
                try {
                    Thread.sleep(2000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            });
        }
    }

    /**
     * newSingleThreadExecutor:
     *
     * 底层：FinalizableDelegatedExecutorService包装的ThreadPoolExecutor实例，corePoolSize为1；maximumPoolSize为1；
     * keepAliveTime为0L；时间单位TimeUnit.MILLISECONDS；workQueue为：
     * new LinkedBlockingQueue<Runnable>() 无解阻塞队列
     * 通俗：创建只有一个线程的线程池，当该线程正繁忙时，对于新任务会进入阻塞队列中(无界的阻塞队列)
     * 适用：按顺序执行任务的场景
     * 创建一个单线程化的线程池，它只会用唯一的工作线程来执行任务，保证所有任务按照指定顺序(FIFO, LIFO, 优先级)执行
     */
    public static void singleTheadPoolTest() {
        ExecutorService pool = Executors.newSingleThreadExecutor();
        for (int i = 0; i < 10; i++) {
            final int ii = i;
            pool.execute(() ->System.out.println(Thread.currentThread().getName() + "=>" + ii));
        }
    }

    /**
     * NewScheduledThreadPool:
     *
     * 底层：创建ScheduledThreadPoolExecutor实例，该对象继承了ThreadPoolExecutor，
     * corePoolSize为传递来的参数，maximumPoolSize为Integer.MAX_VALUE；keepAliveTime为0；
     * 时间单位TimeUnit.NANOSECONDS；workQueue为：new DelayedWorkQueue() 一个按超时时间升序排序的队列
     *
     * 通俗：创建一个固定大小的线程池，线程池内线程存活时间无限制，线程池可以支持定时及周期性任务执行，如果所有线程均处于繁忙状态，对于新任务会进入DelayedWorkQueue队列中，这是一种按照超时时间排序的队列结构
     * 适用：执行周期性任务
     * 创建一个定长线程池，支持定时及周期性任务执行。延迟执行
     */
    public static void scheduleThreadPool() {
        ScheduledExecutorService scheduledThreadPool = Executors.newScheduledThreadPool(5);

        Runnable r1 = new Runnable() {
            @Override
            public void run() {
               System.out.println("线程名称：" + Thread.currentThread().getName() + "，执行:3秒后执行");
            }
        };
        scheduledThreadPool.schedule(r1, 3, TimeUnit.SECONDS);

        Runnable r2 = new Runnable() {
            @Override
            public void run() {
               System.out.println("线程名称：" + Thread.currentThread().getName() + "，执行:延迟2秒后每3秒执行一次");
            }
        };
        scheduledThreadPool.scheduleAtFixedRate(r2, 2, 3, TimeUnit.SECONDS);

        Runnable r3 = new Runnable() {
            @Override
            public void run() {
               System.out.println("线程名称：" + Thread.currentThread().getName() + "，执行:普通任务");
            }
        };
        for (int i = 0; i < 5; i++) {
            scheduledThreadPool.execute(r3);
        }
    }
}

```


