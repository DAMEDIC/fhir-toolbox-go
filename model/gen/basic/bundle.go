package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
)

// A container for a collection of resources.
type Bundle struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A persistent identifier for the bundle that won't change as a bundle is copied from server to server.
	Identifier *Identifier
	// Indicates the purpose of this bundle - how it is intended to be used.
	Type Code
	// The date/time that the bundle was assembled - i.e. when the resources were placed in the bundle.
	Timestamp *Instant
	// If a set of search matches, this is the total number of entries of type 'match' across all pages in the search.  It does not include search.mode = 'include' or 'outcome' entries and it does not provide a count of the number of entries in the Bundle.
	Total *UnsignedInt
	// A series of links that provide context to this bundle.
	Link []BundleLink
	// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
	Entry []BundleEntry
	// Digital Signature - base64 encoded. XML-DSig or a JWT.
	Signature *Signature
}

// A series of links that provide context to this bundle.
type BundleLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A name which details the functional use for this link - see [http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	Relation String
	// The reference details for the link.
	Url Uri
}

// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
type BundleEntry struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A series of links that provide context to this entry.
	Link []BundleLink
	// The Absolute URL for the resource.  The fullUrl SHALL NOT disagree with the id in the resource - i.e. if the fullUrl is not a urn:uuid, the URL shall be version-independent URL consistent with the Resource.id. The fullUrl is a version independent reference to the resource. The fullUrl element SHALL have a value except that:
	// * fullUrl can be empty on a POST (although it does not need to when specifying a temporary id for reference in the bundle)
	// * Results from operations might involve resources that are not identified.
	FullUrl *Uri
	// The Resource for the entry. The purpose/meaning of the resource is determined by the Bundle.type.
	Resource model.Resource
	// Information about the search process that lead to the creation of this entry.
	Search *BundleEntrySearch
	// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
	Request *BundleEntryRequest
	// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
	Response *BundleEntryResponse
}

// Information about the search process that lead to the creation of this entry.
type BundleEntrySearch struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Why this entry is in the result set - whether it's included as a match or because of an _include requirement, or to convey information or warning information about the search process.
	Mode *Code
	// When searching, the server's search ranking score for the entry.
	Score *Decimal
}

// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
type BundleEntryRequest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// In a transaction or batch, this is the HTTP action to be executed for this entry. In a history bundle, this indicates the HTTP action that occurred.
	Method Code
	// The URL for this entry, relative to the root (the address to which the request is posted).
	Url Uri
	// If the ETag values match, return a 304 Not Modified status. See the API documentation for ["Conditional Read"](http.html#cread).
	IfNoneMatch *String
	// Only perform the operation if the last updated date matches. See the API documentation for ["Conditional Read"](http.html#cread).
	IfModifiedSince *Instant
	// Only perform the operation if the Etag value matches. For more information, see the API section ["Managing Resource Contention"](http.html#concurrency).
	IfMatch *String
	// Instruct the server not to perform the create if a specified resource already exists. For further information, see the API documentation for ["Conditional Create"](http.html#ccreate). This is just the query portion of the URL - what follows the "?" (not including the "?").
	IfNoneExist *String
}

// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
type BundleEntryResponse struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The status code returned by processing this entry. The status SHALL start with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP description associated with the status code.
	Status String
	// The location header created by processing this operation, populated if the operation returns a location.
	Location *Uri
	// The Etag for the resource, if the operation for the entry produced a versioned resource (see [Resource Metadata and Versioning](http.html#versioning) and [Managing Resource Contention](http.html#concurrency)).
	Etag *String
	// The date/time that the resource was modified on the server.
	LastModified *Instant
	// An OperationOutcome containing hints and warnings produced as part of processing this entry in a batch or transaction.
	Outcome model.Resource
}

func (r Bundle) ResourceType() string {
	return "Bundle"
}
func (r Bundle) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Bundle) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Identifier != nil {
		s += r.Identifier.MemSize()
	}
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Timestamp != nil {
		s += r.Timestamp.MemSize()
	}
	if r.Total != nil {
		s += r.Total.MemSize()
	}
	for _, i := range r.Link {
		s += i.MemSize()
	}
	s += (cap(r.Link) - len(r.Link)) * int(unsafe.Sizeof(BundleLink{}))
	for _, i := range r.Entry {
		s += i.MemSize()
	}
	s += (cap(r.Entry) - len(r.Entry)) * int(unsafe.Sizeof(BundleEntry{}))
	if r.Signature != nil {
		s += r.Signature.MemSize()
	}
	return s
}
func (r BundleLink) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Relation.MemSize() - int(unsafe.Sizeof(r.Relation))
	s += r.Url.MemSize() - int(unsafe.Sizeof(r.Url))
	return s
}
func (r BundleEntry) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Link {
		s += i.MemSize()
	}
	s += (cap(r.Link) - len(r.Link)) * int(unsafe.Sizeof(BundleLink{}))
	if r.FullUrl != nil {
		s += r.FullUrl.MemSize()
	}
	if r.Resource != nil {
		s += (r.Resource).MemSize()
	}
	if r.Search != nil {
		s += r.Search.MemSize()
	}
	if r.Request != nil {
		s += r.Request.MemSize()
	}
	if r.Response != nil {
		s += r.Response.MemSize()
	}
	return s
}
func (r BundleEntrySearch) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Mode != nil {
		s += r.Mode.MemSize()
	}
	if r.Score != nil {
		s += r.Score.MemSize()
	}
	return s
}
func (r BundleEntryRequest) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Method.MemSize() - int(unsafe.Sizeof(r.Method))
	s += r.Url.MemSize() - int(unsafe.Sizeof(r.Url))
	if r.IfNoneMatch != nil {
		s += r.IfNoneMatch.MemSize()
	}
	if r.IfModifiedSince != nil {
		s += r.IfModifiedSince.MemSize()
	}
	if r.IfMatch != nil {
		s += r.IfMatch.MemSize()
	}
	if r.IfNoneExist != nil {
		s += r.IfNoneExist.MemSize()
	}
	return s
}
func (r BundleEntryResponse) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	if r.Location != nil {
		s += r.Location.MemSize()
	}
	if r.Etag != nil {
		s += r.Etag.MemSize()
	}
	if r.LastModified != nil {
		s += r.LastModified.MemSize()
	}
	if r.Outcome != nil {
		s += (r.Outcome).MemSize()
	}
	return s
}
func (r Bundle) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BundleLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BundleEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BundleEntrySearch) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BundleEntryRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BundleEntryResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Bundle) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Bundle) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Bundle\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Identifier != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		err = r.Identifier.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Type)
		if err != nil {
			return err
		}
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		p := primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_type\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Timestamp != nil && r.Timestamp.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timestamp\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Timestamp)
		if err != nil {
			return err
		}
	}
	if r.Timestamp != nil && (r.Timestamp.Id != nil || r.Timestamp.Extension != nil) {
		p := primitiveElement{Id: r.Timestamp.Id, Extension: r.Timestamp.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_timestamp\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Total != nil && r.Total.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"total\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Total)
		if err != nil {
			return err
		}
	}
	if r.Total != nil && (r.Total.Id != nil || r.Total.Extension != nil) {
		p := primitiveElement{Id: r.Total.Id, Extension: r.Total.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_total\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Link) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"link\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Link {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Entry) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"entry\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Entry {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Signature != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"signature\":"))
		if err != nil {
			return err
		}
		err = r.Signature.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BundleLink) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BundleLink) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Relation)
		if err != nil {
			return err
		}
	}
	if r.Relation.Id != nil || r.Relation.Extension != nil {
		p := primitiveElement{Id: r.Relation.Id, Extension: r.Relation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_relation\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntry) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BundleEntry) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Link) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"link\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Link {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.FullUrl != nil && r.FullUrl.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fullUrl\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.FullUrl)
		if err != nil {
			return err
		}
	}
	if r.FullUrl != nil && (r.FullUrl.Id != nil || r.FullUrl.Extension != nil) {
		p := primitiveElement{Id: r.FullUrl.Id, Extension: r.FullUrl.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_fullUrl\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Resource != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"resource\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Resource)
		if err != nil {
			return err
		}
	}
	if r.Search != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"search\":"))
		if err != nil {
			return err
		}
		err = r.Search.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Request != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"request\":"))
		if err != nil {
			return err
		}
		err = r.Request.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Response != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"response\":"))
		if err != nil {
			return err
		}
		err = r.Response.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntrySearch) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BundleEntrySearch) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Mode != nil && r.Mode.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"mode\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Mode)
		if err != nil {
			return err
		}
	}
	if r.Mode != nil && (r.Mode.Id != nil || r.Mode.Extension != nil) {
		p := primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_mode\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Score != nil && r.Score.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"score\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Score)
		if err != nil {
			return err
		}
	}
	if r.Score != nil && (r.Score.Id != nil || r.Score.Extension != nil) {
		p := primitiveElement{Id: r.Score.Id, Extension: r.Score.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_score\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntryRequest) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BundleEntryRequest) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"method\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Method)
		if err != nil {
			return err
		}
	}
	if r.Method.Id != nil || r.Method.Extension != nil {
		p := primitiveElement{Id: r.Method.Id, Extension: r.Method.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_method\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IfNoneMatch != nil && r.IfNoneMatch.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ifNoneMatch\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IfNoneMatch)
		if err != nil {
			return err
		}
	}
	if r.IfNoneMatch != nil && (r.IfNoneMatch.Id != nil || r.IfNoneMatch.Extension != nil) {
		p := primitiveElement{Id: r.IfNoneMatch.Id, Extension: r.IfNoneMatch.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_ifNoneMatch\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IfModifiedSince != nil && r.IfModifiedSince.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ifModifiedSince\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IfModifiedSince)
		if err != nil {
			return err
		}
	}
	if r.IfModifiedSince != nil && (r.IfModifiedSince.Id != nil || r.IfModifiedSince.Extension != nil) {
		p := primitiveElement{Id: r.IfModifiedSince.Id, Extension: r.IfModifiedSince.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_ifModifiedSince\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IfMatch != nil && r.IfMatch.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ifMatch\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IfMatch)
		if err != nil {
			return err
		}
	}
	if r.IfMatch != nil && (r.IfMatch.Id != nil || r.IfMatch.Extension != nil) {
		p := primitiveElement{Id: r.IfMatch.Id, Extension: r.IfMatch.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_ifMatch\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IfNoneExist != nil && r.IfNoneExist.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ifNoneExist\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IfNoneExist)
		if err != nil {
			return err
		}
	}
	if r.IfNoneExist != nil && (r.IfNoneExist.Id != nil || r.IfNoneExist.Extension != nil) {
		p := primitiveElement{Id: r.IfNoneExist.Id, Extension: r.IfNoneExist.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_ifNoneExist\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntryResponse) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BundleEntryResponse) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Location != nil && r.Location.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"location\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Location)
		if err != nil {
			return err
		}
	}
	if r.Location != nil && (r.Location.Id != nil || r.Location.Extension != nil) {
		p := primitiveElement{Id: r.Location.Id, Extension: r.Location.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_location\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Etag != nil && r.Etag.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"etag\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Etag)
		if err != nil {
			return err
		}
	}
	if r.Etag != nil && (r.Etag.Id != nil || r.Etag.Extension != nil) {
		p := primitiveElement{Id: r.Etag.Id, Extension: r.Etag.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_etag\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastModified != nil && r.LastModified.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastModified\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastModified)
		if err != nil {
			return err
		}
	}
	if r.LastModified != nil && (r.LastModified.Id != nil || r.LastModified.Extension != nil) {
		p := primitiveElement{Id: r.LastModified.Id, Extension: r.LastModified.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastModified\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Outcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"outcome\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Outcome)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r Bundle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Bundle"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Timestamp, xml.StartElement{Name: xml.Name{Local: "timestamp"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Total, xml.StartElement{Name: xml.Name{Local: "total"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Entry, xml.StartElement{Name: xml.Name{Local: "entry"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Signature, xml.StartElement{Name: xml.Name{Local: "signature"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BundleLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relation, xml.StartElement{Name: xml.Name{Local: "relation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FullUrl, xml.StartElement{Name: xml.Name{Local: "fullUrl"}})
	if err != nil {
		return err
	}
	if r.Resource != nil {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "resource"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "resource"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Search, xml.StartElement{Name: xml.Name{Local: "search"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Response, xml.StartElement{Name: xml.Name{Local: "response"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntrySearch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Score, xml.StartElement{Name: xml.Name{Local: "score"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntryRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfNoneMatch, xml.StartElement{Name: xml.Name{Local: "ifNoneMatch"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfModifiedSince, xml.StartElement{Name: xml.Name{Local: "ifModifiedSince"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfMatch, xml.StartElement{Name: xml.Name{Local: "ifMatch"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfNoneExist, xml.StartElement{Name: xml.Name{Local: "ifNoneExist"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BundleEntryResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Etag, xml.StartElement{Name: xml.Name{Local: "etag"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastModified, xml.StartElement{Name: xml.Name{Local: "lastModified"}})
	if err != nil {
		return err
	}
	if r.Outcome != nil {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "outcome"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "outcome"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r Bundle) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		if r.Identifier != nil {
			children = append(children, *r.Identifier)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "timestamp") {
		if r.Timestamp != nil {
			children = append(children, *r.Timestamp)
		}
	}
	if len(name) == 0 || slices.Contains(name, "total") {
		if r.Total != nil {
			children = append(children, *r.Total)
		}
	}
	if len(name) == 0 || slices.Contains(name, "link") {
		for _, v := range r.Link {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "entry") {
		for _, v := range r.Entry {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "signature") {
		if r.Signature != nil {
			children = append(children, *r.Signature)
		}
	}
	return children
}
func (r Bundle) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Bundle to Boolean")
}
func (r Bundle) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Bundle to String")
}
func (r Bundle) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Bundle to Integer")
}
func (r Bundle) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Bundle to Decimal")
}
func (r Bundle) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Bundle to Date")
}
func (r Bundle) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Bundle to Time")
}
func (r Bundle) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Bundle to DateTime")
}
func (r Bundle) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Bundle to Quantity")
}
func (r Bundle) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Bundle
	switch other := other.(type) {
	case Bundle:
		o = &other
	case *Bundle:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r Bundle) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(Bundle)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r Bundle) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Timestamp",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "Total",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "UnsignedInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Link",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BundleLink",
				Namespace: "FHIR",
			},
		}, {
			Name: "Entry",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BundleEntry",
				Namespace: "FHIR",
			},
		}, {
			Name: "Signature",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Signature",
				Namespace: "FHIR",
			},
		}},
		Name:      "Bundle",
		Namespace: "FHIR",
	}
}
func (r BundleLink) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relation") {
		children = append(children, r.Relation)
	}
	if len(name) == 0 || slices.Contains(name, "url") {
		children = append(children, r.Url)
	}
	return children
}
func (r BundleLink) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BundleLink to Boolean")
}
func (r BundleLink) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BundleLink to String")
}
func (r BundleLink) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BundleLink to Integer")
}
func (r BundleLink) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BundleLink to Decimal")
}
func (r BundleLink) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BundleLink to Date")
}
func (r BundleLink) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BundleLink to Time")
}
func (r BundleLink) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BundleLink to DateTime")
}
func (r BundleLink) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BundleLink to Quantity")
}
func (r BundleLink) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BundleLink
	switch other := other.(type) {
	case BundleLink:
		o = &other
	case *BundleLink:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r BundleLink) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(BundleLink)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r BundleLink) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Relation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}},
		Name:      "BundleLink",
		Namespace: "FHIR",
	}
}
func (r BundleEntry) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "link") {
		for _, v := range r.Link {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fullUrl") {
		if r.FullUrl != nil {
			children = append(children, *r.FullUrl)
		}
	}
	if len(name) == 0 || slices.Contains(name, "resource") {
		if r.Resource != nil {
			children = append(children, r.Resource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "search") {
		if r.Search != nil {
			children = append(children, *r.Search)
		}
	}
	if len(name) == 0 || slices.Contains(name, "request") {
		if r.Request != nil {
			children = append(children, *r.Request)
		}
	}
	if len(name) == 0 || slices.Contains(name, "response") {
		if r.Response != nil {
			children = append(children, *r.Response)
		}
	}
	return children
}
func (r BundleEntry) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BundleEntry to Boolean")
}
func (r BundleEntry) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BundleEntry to String")
}
func (r BundleEntry) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BundleEntry to Integer")
}
func (r BundleEntry) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BundleEntry to Decimal")
}
func (r BundleEntry) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BundleEntry to Date")
}
func (r BundleEntry) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BundleEntry to Time")
}
func (r BundleEntry) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BundleEntry to DateTime")
}
func (r BundleEntry) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BundleEntry to Quantity")
}
func (r BundleEntry) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BundleEntry
	switch other := other.(type) {
	case BundleEntry:
		o = &other
	case *BundleEntry:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r BundleEntry) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(BundleEntry)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r BundleEntry) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Link",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BundleLink",
				Namespace: "FHIR",
			},
		}, {
			Name: "FullUrl",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Resource",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Search",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "BundleEntrySearch",
				Namespace: "FHIR",
			},
		}, {
			Name: "Request",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "BundleEntryRequest",
				Namespace: "FHIR",
			},
		}, {
			Name: "Response",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "BundleEntryResponse",
				Namespace: "FHIR",
			},
		}},
		Name:      "BundleEntry",
		Namespace: "FHIR",
	}
}
func (r BundleEntrySearch) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "mode") {
		if r.Mode != nil {
			children = append(children, *r.Mode)
		}
	}
	if len(name) == 0 || slices.Contains(name, "score") {
		if r.Score != nil {
			children = append(children, *r.Score)
		}
	}
	return children
}
func (r BundleEntrySearch) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BundleEntrySearch to Boolean")
}
func (r BundleEntrySearch) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BundleEntrySearch to String")
}
func (r BundleEntrySearch) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BundleEntrySearch to Integer")
}
func (r BundleEntrySearch) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BundleEntrySearch to Decimal")
}
func (r BundleEntrySearch) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BundleEntrySearch to Date")
}
func (r BundleEntrySearch) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BundleEntrySearch to Time")
}
func (r BundleEntrySearch) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BundleEntrySearch to DateTime")
}
func (r BundleEntrySearch) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BundleEntrySearch to Quantity")
}
func (r BundleEntrySearch) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BundleEntrySearch
	switch other := other.(type) {
	case BundleEntrySearch:
		o = &other
	case *BundleEntrySearch:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r BundleEntrySearch) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(BundleEntrySearch)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r BundleEntrySearch) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Mode",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Score",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}},
		Name:      "BundleEntrySearch",
		Namespace: "FHIR",
	}
}
func (r BundleEntryRequest) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "method") {
		children = append(children, r.Method)
	}
	if len(name) == 0 || slices.Contains(name, "url") {
		children = append(children, r.Url)
	}
	if len(name) == 0 || slices.Contains(name, "ifNoneMatch") {
		if r.IfNoneMatch != nil {
			children = append(children, *r.IfNoneMatch)
		}
	}
	if len(name) == 0 || slices.Contains(name, "ifModifiedSince") {
		if r.IfModifiedSince != nil {
			children = append(children, *r.IfModifiedSince)
		}
	}
	if len(name) == 0 || slices.Contains(name, "ifMatch") {
		if r.IfMatch != nil {
			children = append(children, *r.IfMatch)
		}
	}
	if len(name) == 0 || slices.Contains(name, "ifNoneExist") {
		if r.IfNoneExist != nil {
			children = append(children, *r.IfNoneExist)
		}
	}
	return children
}
func (r BundleEntryRequest) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BundleEntryRequest to Boolean")
}
func (r BundleEntryRequest) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BundleEntryRequest to String")
}
func (r BundleEntryRequest) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BundleEntryRequest to Integer")
}
func (r BundleEntryRequest) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BundleEntryRequest to Decimal")
}
func (r BundleEntryRequest) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BundleEntryRequest to Date")
}
func (r BundleEntryRequest) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BundleEntryRequest to Time")
}
func (r BundleEntryRequest) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BundleEntryRequest to DateTime")
}
func (r BundleEntryRequest) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BundleEntryRequest to Quantity")
}
func (r BundleEntryRequest) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BundleEntryRequest
	switch other := other.(type) {
	case BundleEntryRequest:
		o = &other
	case *BundleEntryRequest:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r BundleEntryRequest) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(BundleEntryRequest)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r BundleEntryRequest) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Method",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "IfNoneMatch",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "IfModifiedSince",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "IfMatch",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "IfNoneExist",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		Name:      "BundleEntryRequest",
		Namespace: "FHIR",
	}
}
func (r BundleEntryResponse) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "location") {
		if r.Location != nil {
			children = append(children, *r.Location)
		}
	}
	if len(name) == 0 || slices.Contains(name, "etag") {
		if r.Etag != nil {
			children = append(children, *r.Etag)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lastModified") {
		if r.LastModified != nil {
			children = append(children, *r.LastModified)
		}
	}
	if len(name) == 0 || slices.Contains(name, "outcome") {
		if r.Outcome != nil {
			children = append(children, r.Outcome)
		}
	}
	return children
}
func (r BundleEntryResponse) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BundleEntryResponse to Boolean")
}
func (r BundleEntryResponse) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BundleEntryResponse to String")
}
func (r BundleEntryResponse) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BundleEntryResponse to Integer")
}
func (r BundleEntryResponse) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BundleEntryResponse to Decimal")
}
func (r BundleEntryResponse) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BundleEntryResponse to Date")
}
func (r BundleEntryResponse) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BundleEntryResponse to Time")
}
func (r BundleEntryResponse) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BundleEntryResponse to DateTime")
}
func (r BundleEntryResponse) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BundleEntryResponse to Quantity")
}
func (r BundleEntryResponse) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BundleEntryResponse
	switch other := other.(type) {
	case BundleEntryResponse:
		o = &other
	case *BundleEntryResponse:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r BundleEntryResponse) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(BundleEntryResponse)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r BundleEntryResponse) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Location",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Etag",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastModified",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "Outcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "",
				Namespace: "FHIR",
			},
		}},
		Name:      "BundleEntryResponse",
		Namespace: "FHIR",
	}
}
