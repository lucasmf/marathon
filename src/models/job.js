// marathon
// https://github.com/topfreegames/marathon
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2016 Top Free Games <backend@tfgco.com>

import Sequelize from 'sequelize'

module.exports = sequelize => (
  sequelize.define('jobs', {
    id: {
      type: Sequelize.UUID,
      primaryKey: true,
      defaultValue: Sequelize.UUIDV4,
    },
    totalBatches: {
      type: Sequelize.INTEGER,
      validate: { min: 1 },
      field: 'total_batches',
    },
    completedBatches: {
      type: Sequelize.INTEGER,
      allowNull: false,
      validate: { min: 0 },
      defaultValue: 0,
      field: 'completed_batches',
    },
    completedAt: {
      type: Sequelize.DATE,
      defaultValue: Sequelize.fn('now'),
      field: 'completed_at',
    },
    expireAt: {
      type: Sequelize.DATE,
      field: 'expire_at',
    },
    context: {
      type: Sequelize.JSONB,
      allowNull: false,
    },
    service: {
      type: Sequelize.ENUM,
      allowNull: false,
      values: ['apns', 'gcm'],
    },
    filters: {
      type: Sequelize.JSONB,
    },
    csvUrl: {
      type: Sequelize.STRING,
      validate: { isUrl: true },
      field: 'csv_url',
    },
    createdBy: {
      type: Sequelize.STRING,
      allowNull: false,
      validate: { len: [1, 2000] },
      field: 'created_by',
    },
  }, {
    timestamps: true,
    underscored: true,
    classMethods: {
      associate: (models) => {
        models.Job.belongsTo(models.App, {
          foreignKey: {
            allowNull: false,
            field: 'app_id',
            fieldName: 'appId',
          },
        })
        models.Job.belongsTo(models.Template, {
          foreignKey: {
            allowNull: false,
            field: 'template_id',
            fieldName: 'templateId',
          },
        })
      },
    },
  })
)