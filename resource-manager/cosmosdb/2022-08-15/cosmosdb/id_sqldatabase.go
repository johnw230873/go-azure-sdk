package cosmosdb

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SqlDatabaseId{}

// SqlDatabaseId is a struct representing the Resource ID for a Sql Database
type SqlDatabaseId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	DatabaseName      string
}

// NewSqlDatabaseID returns a new SqlDatabaseId struct
func NewSqlDatabaseID(subscriptionId string, resourceGroupName string, accountName string, databaseName string) SqlDatabaseId {
	return SqlDatabaseId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		DatabaseName:      databaseName,
	}
}

// ParseSqlDatabaseID parses 'input' into a SqlDatabaseId
func ParseSqlDatabaseID(input string) (*SqlDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, fmt.Errorf("the segment 'databaseName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSqlDatabaseIDInsensitively parses 'input' case-insensitively into a SqlDatabaseId
// note: this method should only be used for API response data and not user input
func ParseSqlDatabaseIDInsensitively(input string) (*SqlDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, fmt.Errorf("the segment 'databaseName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSqlDatabaseID checks that 'input' can be parsed as a Sql Database ID
func ValidateSqlDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Database ID
func (id SqlDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.DatabaseName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Database ID
func (id SqlDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticSqlDatabases", "sqlDatabases", "sqlDatabases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
	}
}

// String returns a human-readable description of this Sql Database ID
func (id SqlDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
	}
	return fmt.Sprintf("Sql Database (%s)", strings.Join(components, "\n"))
}
