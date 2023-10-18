package job

import (
	"testing"

	"github.com/buildkite/agent/v3/internal/job/shell"
	"github.com/stretchr/testify/assert"
)

func makeExecutor(t *testing.T) *Executor {
	var err error
	executor := New(ExecutorConfig{
		Plugins:        "[{\"some-org/some-repo\": {}}]",
		PluginsEnabled: true,
	})

	executor.shell, err = shell.New()
	assert.NoError(t, err)

	return executor
}

func TestPreparePlugins_NoAllowlistConfigured(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	err := executor.preparePlugins()
	assert.NoError(t, err)
}

func TestPreparePlugins_WithEmptyAllowlist(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	executor.PluginAllowlist = make([]string, 0)
	err := executor.preparePlugins()
	assert.Error(t, err)
	assert.ErrorContains(t, err, "allowlist is empty")
}

func TestPreparePlugins_MatchingAllowlist(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	executor.PluginAllowlist = []string{"some-org/some-repo"}
	err := executor.preparePlugins()
	assert.NoError(t, err)
}

func TestPreparePlugins_FailingAllowlist(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	executor.PluginAllowlist = []string{"some-other-org/some-repo"}
	err := executor.preparePlugins()
	assert.ErrorContains(t, err, "Failed plugin allowlist")
	assert.ErrorContains(t, err, "some-org/some-repo")
}

func TestPreparePlugins_VendoredPluginAllowed(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	executor.Plugins = "[{\"./some/local/file\": {}}]"
	executor.PluginAllowlist = []string{"some-other-org/some-repo"}
	err := executor.preparePlugins()
	assert.NoError(t, err)
}

func TestPreparePlugins_LocalFileAllowed(t *testing.T) {
	t.Parallel()

	executor := makeExecutor(t)
	executor.Plugins = "[{\"file:///some/agent/file\": {}}]"
	executor.PluginAllowlist = []string{"some-other-org/some-repo"}
	err := executor.preparePlugins()
	assert.NoError(t, err)
}
