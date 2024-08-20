import asyncio
import time

async def heavy_computation(id, result_queue):
    # Simulate heavy computation by calculating sum of squares for a large range
    # print("innn")
    total = sum(i**2 for i in range(100000))
    await result_queue.put(total)

async def main():
    start_time = time.time()  # Start measuring time

    result_queue = asyncio.Queue()

    tasks = []
    for i in range(500):
        tasks.append(heavy_computation(i, result_queue))
    await asyncio.gather(*tasks)

    total_result = 0
    while not result_queue.empty():
        total_result += await result_queue.get()

    elapsed_time = time.time() - start_time  # Calculate elapsed time
    print("Total time:", elapsed_time)
    print("Total result:", total_result)

asyncio.run(main())
