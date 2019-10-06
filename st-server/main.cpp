/*
https://github.com/winlinvip/study-coroutine#usage
*/
#include <stdio.h>
#include <string.h>
#include <assert.h>
#include <st.h>
#include <arpa/inet.h>

#define LISTEN_PORT 8080

void* serve(void* arg)
{
    printf("start serve client\n");
    st_netfd_t fd = (st_netfd_t)arg;
    char buf[1024];
    ssize_t nn_msgs = 0;
    for (;;) {
        int nn = st_read(fd, buf, sizeof(buf), ST_UTIME_NO_TIMEOUT);
        if (nn <= 0) {
            break;
        }

        if (st_write(fd, buf, nn, ST_UTIME_NO_TIMEOUT) <= 0) {
            break;
        }

        nn_msgs++;
    }

    printf("server done, msgs=%d\n", nn_msgs);
    return NULL;
}

int main(int argc, char** argv)
{
    assert(st_set_eventsys(ST_EVENTSYS_ALT) >= 0); // Use epoll.
    assert(st_init() >= 0);

    // Server listen at 8080
    int sock = socket(PF_INET, SOCK_STREAM, 0);
    assert(sock > 0);

    int n = 1;
    assert(setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, (char *)&n, sizeof(n)) >= 0);

    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(LISTEN_PORT);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);
    assert(bind(sock, (struct sockaddr *)&addr, sizeof(addr)) >= 0);

    assert(listen(sock, 10) >= 0);

    st_netfd_t fd = st_netfd_open_socket(sock);
    assert(fd != NULL);

    printf("st_getfdlimit=%d, listen=%d\n", st_getfdlimit(), LISTEN_PORT);

    for (;;) {
        st_netfd_t cfd = st_accept(fd, NULL, NULL, ST_UTIME_NO_TIMEOUT);
        st_thread_create(serve, cfd, 0, 0);
    }

    return 0;
}
