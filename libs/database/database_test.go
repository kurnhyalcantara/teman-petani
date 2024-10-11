package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/kurnhyalcantara/teman-petani/libs/database/wrapper/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// DatabaseTestSuite adalah struktur yang menggunakan testify suite
type DatabaseTestSuite struct {
	suite.Suite
	db          *DB
	mockWrapper *mocks.DatabaseConnectionInterface
}

// SetupTest dijalankan sebelum setiap test
func (s *DatabaseTestSuite) SetupTest() {
	s.mockWrapper = new(mocks.DatabaseConnectionInterface)
	s.db = &DB{
		DriverName: "postgres",
		Config:     &Config{},
	}
}

// TestInitDatabase menguji inisialisasi database
func TestInitDatabase(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}

// TestConnect menguji metode Connect
func (s *DatabaseTestSuite) TestConnect() {
	s.Run("Success", func() {
		s.mockWrapper.On("Open", "postgres", mock.Anything).Return(&sql.DB{}, nil)

		err := s.db.Connect()
		assert.NoError(s.T(), err)
		s.mockWrapper.AssertExpectations(s.T())
	})

	s.Run("Failure", func() {
		s.mockWrapper.On("Open", "postgres", mock.Anything).Return(nil, errors.New("connection error"))

		err := s.db.Connect()
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "connection error", err.Error())
		s.mockWrapper.AssertExpectations(s.T())
	})
}

// TestTryConnect menguji metode TryConnect
func (s *DatabaseTestSuite) TestTryConnect() {
	s.Run("Success", func() {
		s.mockWrapper.On("Open", "postgres", mock.Anything).Return(&sql.DB{}, nil).Once()

		err := s.db.TryConnect()
		assert.NoError(s.T(), err)
		s.mockWrapper.AssertExpectations(s.T())
	})

	s.Run("MaxRetryExceeded", func() {
		s.mockWrapper.On("Open", "postgres", mock.Anything).Return(nil, errors.New("connection error")).Times(3)

		err := s.db.TryConnect()
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "connection error", err.Error())
		s.mockWrapper.AssertExpectations(s.T())
	})
}

// TestCheckConnection menguji metode CheckConnection
func (s *DatabaseTestSuite) TestCheckConnection() {
	s.Run("Success", func() {
		s.mockWrapper.On("Ping").Return(nil)

		err := s.db.CheckConnection()
		assert.NoError(s.T(), err)
		s.mockWrapper.AssertExpectations(s.T())
	})

	s.Run("PingFailure", func() {
		s.mockWrapper.On("Ping").Return(errors.New("ping error"))
		s.mockWrapper.On("Close").Return(nil)
		s.mockWrapper.On("Open", "postgres", mock.Anything).Return(&sql.DB{}, nil)

		err := s.db.CheckConnection()
		assert.NoError(s.T(), err)
		s.mockWrapper.AssertExpectations(s.T())
	})
}

// TestStartTransaction menguji metode StartTransaction
func (s *DatabaseTestSuite) TestStartTransaction() {
	s.Run("Success", func() {
		s.mockWrapper.On("Begin").Return(&sql.Tx{}, nil)

		err := s.db.StartTransaction()
		assert.NoError(s.T(), err)
		s.mockWrapper.AssertExpectations(s.T())
	})

	s.Run("Failure", func() {
		s.mockWrapper.On("Begin").Return(nil, errors.New("transaction error"))

		err := s.db.StartTransaction()
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "transaction error", err.Error())
		s.mockWrapper.AssertExpectations(s.T())
	})
}

// TestAddCounter menguji metode AddCounter
func (s *DatabaseTestSuite) TestAddCounter() {
	s.db.AddCounter()
	assert.Equal(s.T(), 1, s.db.Counter)

	s.db.AddCounter()
	assert.Equal(s.T(), 2, s.db.Counter)
}

// TestSetMaxIdleConnections menguji metode SetMaxIdleConnections
func (s *DatabaseTestSuite) TestSetMaxIdleConnections() {
	s.Run("SetMaxIdleConnections", func() {
		s.mockWrapper.On("SetMaxIdleConnections", 5)

		s.db.SetMaxIdleConnections(5)
		s.mockWrapper.AssertExpectations(s.T())
	})

	s.Run("SetMaxOpenConnections", func() {
		s.mockWrapper.On("SetMaxOpenConnections", 10)

		s.db.SetMaxOpenConnections(10)
		s.mockWrapper.AssertExpectations(s.T())
	})
}
