// Code generated by https://github.com/selefra/selefra-terraform-provider-scaffolding DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***
package resources

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

func GetSelefraProvider() *provider.Provider {
	diagnostics := schema.NewDiagnostics()
	selefraProvider, d := GetSelefraTerraformProvider().ToSelefraProvider(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) *bridge.TerraformBridge {
		return taskClient.(*Client).TerraformBridge
	})
	if diagnostics.AddDiagnostics(d).HasError() {
		panic(diagnostics.ToString())
	}

    selefraProvider.TableList = GetSelefraTables()

	return selefraProvider
}

func GetSelefraTables() []*schema.Table {

    diagnostics := schema.NewDiagnostics()
    tables := make([]*schema.Table, 0)
    var table *schema.Table
    var d *schema.Diagnostics

    
    table, d = TableSchemaGenerator_auth0_rule_config()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_role()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_custom_domain_verification()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_action()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_email()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_guardian()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_client_grant()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_hook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_organization_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_resource_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_user()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_prompt()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_organization()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_trigger_binding()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_branding_theme()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_client()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_log_stream()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_organization_member()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_global_client()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_auth0_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    

    if diagnostics.HasError() {
        panic(diagnostics.ToString())
    }

	return tables
}
