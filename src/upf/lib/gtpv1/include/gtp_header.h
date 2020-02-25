#ifndef __GTP_HEADER_H__
#define __GTP_HEADER_H__

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

#include <arpa/inet.h>

#include "utlt_lib.h"

#define GTPV1_HEADER_LEN 8
#define GTPV1_OPT_HEADER_LEN 4

// TODO : Handle ntohl and ntohs for non-uint8_t type
typedef struct _Gtpv1Header {
    struct {
        ENDIAN6(
            uint8_t version:3;,
            uint8_t PT:1;,
            uint8_t spare:1;,
            uint8_t extHdrFlag:1;,
            uint8_t seqFlag:1;,
            uint8_t pn:1;
        )
    };
    
    uint8_t type;
    uint16_t _length;
    uint32_t _teid;
} __attribute__ ((packed)) Gtpv1Header;

typedef struct _Gtp1OptHeader {
    uint16_t _seqNum;
    uint8_t nPdnNum;
    uint8_t nextExtHrdType;

} __attribute__ ((packed)) Gtp1OptHeader;

#define GTPV1_ECHO_REQUEST     1
#define GTPV1_ECHO_RESPONSE    2
#define GTPV1_ERROR_INDICATION 26
#define GTPV1_END_MARK         254
#define GTPV1_T_PDU            255

#ifdef __cplusplus
}
#endif /* __cplusplus */

#endif /* __GTP_HEADER_H__ */
