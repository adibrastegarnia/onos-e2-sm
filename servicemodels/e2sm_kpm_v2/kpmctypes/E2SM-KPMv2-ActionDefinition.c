/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-rm.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "E2SM-KPMv2-ActionDefinition.h"

#include "E2SM-KPMv2-ActionDefinition-Format1.h"
#include "E2SM-KPMv2-ActionDefinition-Format2.h"
#include "E2SM-KPMv2-ActionDefinition-Format3.h"
static asn_per_constraints_t asn_PER_type_actionDefinition_formats_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  2,  2,  0,  2 }	/* (0..2,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_actionDefinition_formats_3[] = {
	{ ATF_POINTER, 0, offsetof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats, choice.actionDefinition_Format1),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_E2SM_KPMv2_ActionDefinition_Format1,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"actionDefinition-Format1"
		},
	{ ATF_POINTER, 0, offsetof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats, choice.actionDefinition_Format2),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_E2SM_KPMv2_ActionDefinition_Format2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"actionDefinition-Format2"
		},
	{ ATF_POINTER, 0, offsetof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats, choice.actionDefinition_Format3),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_E2SM_KPMv2_ActionDefinition_Format3,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"actionDefinition-Format3"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_actionDefinition_formats_tag2el_3[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* actionDefinition-Format1 */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* actionDefinition-Format2 */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* actionDefinition-Format3 */
};
static asn_CHOICE_specifics_t asn_SPC_actionDefinition_formats_specs_3 = {
	sizeof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats),
	offsetof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats, _asn_ctx),
	offsetof(struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats, present),
	sizeof(((struct E2SM_KPMv2_ActionDefinition__actionDefinition_formats *)0)->present),
	asn_MAP_actionDefinition_formats_tag2el_3,
	3,	/* Count of tags in the map */
	0, 0,
	3	/* Extensions start */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_actionDefinition_formats_3 = {
	"actionDefinition-formats",
	"actionDefinition-formats",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ 0, &asn_PER_type_actionDefinition_formats_constr_3, CHOICE_constraint },
	asn_MBR_actionDefinition_formats_3,
	3,	/* Elements count */
	&asn_SPC_actionDefinition_formats_specs_3	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_E2SM_KPMv2_ActionDefinition_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct E2SM_KPMv2_ActionDefinition, ric_Style_Type),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_RIC_Style_Type,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"ric-Style-Type"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct E2SM_KPMv2_ActionDefinition, actionDefinition_formats),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_actionDefinition_formats_3,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"actionDefinition-formats"
		},
};
static const ber_tlv_tag_t asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_E2SM_KPMv2_ActionDefinition_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* ric-Style-Type */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* actionDefinition-formats */
};
static asn_SEQUENCE_specifics_t asn_SPC_E2SM_KPMv2_ActionDefinition_specs_1 = {
	sizeof(struct E2SM_KPMv2_ActionDefinition),
	offsetof(struct E2SM_KPMv2_ActionDefinition, _asn_ctx),
	asn_MAP_E2SM_KPMv2_ActionDefinition_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	2,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_E2SM_KPMv2_ActionDefinition = {
	"E2SM-KPMv2-ActionDefinition",
	"E2SM-KPMv2-ActionDefinition",
	&asn_OP_SEQUENCE,
	asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1,
	sizeof(asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1[0]), /* 1 */
	asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1,	/* Same as above */
	sizeof(asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_KPMv2_ActionDefinition_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_E2SM_KPMv2_ActionDefinition_1,
	2,	/* Elements count */
	&asn_SPC_E2SM_KPMv2_ActionDefinition_specs_1	/* Additional specs */
};
