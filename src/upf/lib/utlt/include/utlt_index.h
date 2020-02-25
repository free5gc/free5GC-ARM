#ifndef __INDEX_H__
#define __INDEX_H__

#include <pthread.h>

/*****************************************************************************
 * The structure must have variable named "index" when it is in the parameter.
 * Do not modify the variable named "index" after call function "IndexAlloc".
 * How to use the Index can look the test file "testIndex".
 *****************************************************************************/

#define IndexDeclare(__name, __type, __cap) \
    typedef struct { \
        int qFront, qEnd, qCap; \
        int arrayForIndex[__cap + 1]; \
        __type *queueAvail[__cap + 1], pool[__cap + 1]; \
        pthread_mutex_t lock; \
    } index##__name##_t; \
    index##__name##_t __name

#define IndexInit(__nameptr, __cap) do { \
    (__nameptr)->qFront = 0; \
    (__nameptr)->qEnd = __cap; \
    (__nameptr)->qCap = __cap; \
    for (int __i = 0; __i < __cap; __i++) { \
        (__nameptr)->arrayForIndex[__i] = __i; \
        (__nameptr)->queueAvail[__i] = &((__nameptr)->pool[__i]); \
    } \
    pthread_mutex_init(&(__nameptr)->lock, 0); \
} while(0)

#define IndexTerminate(__nameptr) pthread_mutex_destroy(&(__nameptr)->lock)

// The number of available space in this pool
#define IndexSize(__nameptr) \
    ((((__nameptr)->qEnd + (__nameptr)->qCap + 1 - (__nameptr)->qFront)) % ((__nameptr)->qCap + 1))

// Total space of this pool, including used and unused
#define IndexCap(__nameptr) ((__nameptr)->qCap)

#define IndexAlloc(__nameptr, __assignedPtr) do { \
    pthread_mutex_lock(&(__nameptr)->lock); \
    if (IndexSize(__nameptr) > 0) { \
        (__assignedPtr) = (__nameptr)->queueAvail[(__nameptr)->qFront]; \
        (__assignedPtr)->index = (__nameptr)->arrayForIndex[(__nameptr)->qFront]; \
        (__nameptr)->qFront = ((__nameptr)->qFront + 1) % ((__nameptr)->qCap + 1); \
    } else { \
        (__assignedPtr) = NULL; \
    } \
    pthread_mutex_unlock(&(__nameptr)->lock); \
} while(0)

#define IndexFree(__nameptr, __assignedPtr) do { \
    pthread_mutex_lock(&(__nameptr)->lock); \
    if (IndexSize(__nameptr) < (__nameptr)->qCap) { \
        (__nameptr)->queueAvail[(__nameptr)->qEnd] = (__assignedPtr); \
        (__nameptr)->arrayForIndex[(__nameptr)->qEnd] = (__assignedPtr)->index; \
        (__nameptr)->qEnd = ((__nameptr)->qEnd + 1) % ((__nameptr)->qCap + 1); \
    } \
    pthread_mutex_unlock(&(__nameptr)->lock); \
} while(0)

#define IndexFind(__nameptr, __index) (void *) &(__nameptr)->pool[__index];

#endif /* __INDEX_H__ */
