/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_Choice4_H_
#define	_Choice4_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Choice4_PR {
	Choice4_PR_NOTHING,	/* No components present */
	Choice4_PR_choice4A
	/* Extensions may appear below */
	
} Choice4_PR;

/* Choice4 */
typedef struct Choice4 {
	Choice4_PR present;
	union Choice4_u {
		long	 choice4A;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Choice4_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Choice4;
extern asn_CHOICE_specifics_t asn_SPC_Choice4_specs_1;
extern asn_TYPE_member_t asn_MBR_Choice4_1[1];
extern asn_per_constraints_t asn_PER_type_Choice4_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _Choice4_H_ */
#include "asn_internal.h"
